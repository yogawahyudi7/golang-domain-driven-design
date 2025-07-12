package services

import (
	"time"

	"golang-domain-driven-design/internal/application/usecases"
	"golang-domain-driven-design/internal/infrastructure/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTService handles JWT token operations
type JWTService interface {
	GenerateToken(user *usecases.UserResponse) (*usecases.AuthResponse, error)
	ValidateToken(tokenString string) (*JWTClaims, error)
	RefreshToken(tokenString string) (*usecases.AuthResponse, error)
}

// JWTClaims represents the claims structure
type JWTClaims struct {
	UserID   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	IsActive bool      `json:"is_active"`
	jwt.RegisteredClaims
}

// jwtService implements JWTService interface
type jwtService struct {
	secretKey string
	issuer    string
	expiresIn time.Duration
}

// NewJWTService creates a new JWT service
func NewJWTService(cfg *config.Config) JWTService {
	return &jwtService{
		secretKey: cfg.JWT.Secret,
		issuer:    cfg.App.Name,
		expiresIn: cfg.JWT.ExpiresIn,
	}
}

// GenerateToken generates a new JWT token for the user
func (j *jwtService) GenerateToken(user *usecases.UserResponse) (*usecases.AuthResponse, error) {
	now := time.Now()
	expiresAt := now.Add(j.expiresIn)

	claims := &JWTClaims{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Username,
		IsActive: user.IsActive,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    j.issuer,
			Subject:   user.ID.String(),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return nil, err
	}

	return &usecases.AuthResponse{
		User:        user,
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   int64(j.expiresIn.Seconds()),
		ExpiresAt:   expiresAt.Unix(),
	}, nil
}

// ValidateToken validates a JWT token and returns the claims
func (j *jwtService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		// Check if user is still active
		if !claims.IsActive {
			return nil, jwt.ErrTokenInvalidClaims
		}
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}

// RefreshToken generates a new token from an existing valid token
func (j *jwtService) RefreshToken(tokenString string) (*usecases.AuthResponse, error) {
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	// Create a user response from claims for generating new token
	user := &usecases.UserResponse{
		ID:       claims.UserID,
		Email:    claims.Email,
		Username: claims.Username,
		IsActive: claims.IsActive,
	}

	return j.GenerateToken(user)
}

// ExtractTokenFromHeader extracts JWT token from Authorization header
func ExtractTokenFromHeader(authHeader string) string {
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		return authHeader[7:]
	}
	return ""
}
