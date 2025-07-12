package middleware

import (
	"net/http"
	"strings"

	"golang-domain-driven-design/internal/application/services"

	"github.com/gin-gonic/gin"
)

// JWTAuth holds JWT service for authentication middleware
type JWTAuth struct {
	jwtService services.JWTService
}

// NewJWTAuth creates new JWT authentication middleware
func NewJWTAuth(jwtService services.JWTService) *JWTAuth {
	return &JWTAuth{
		jwtService: jwtService,
	}
}

// Middleware returns the gin middleware function
func (j *JWTAuth) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		tokenString := services.ExtractTokenFromHeader(authHeader)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid authorization header format. Use 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		claims, err := j.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_username", claims.Username)
		c.Set("is_active", claims.IsActive)

		c.Next()
	}
}

// OptionalJWTAuthMiddleware creates an optional JWT authentication middleware
// It won't abort the request if no token is provided, but will set user info if valid token exists
func OptionalJWTAuthMiddleware(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		tokenString := services.ExtractTokenFromHeader(authHeader)
		if tokenString == "" {
			c.Next()
			return
		}

		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.Next()
			return
		}

		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_username", claims.Username)
		c.Set("is_active", claims.IsActive)

		c.Next()
	}
}

// AdminMiddleware checks if the user has admin privileges
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// For now, we'll implement a simple check
		// In a real application, you would check user roles from database
		userEmail, exists := c.Get("user_email")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "User not authenticated",
			})
			c.Abort()
			return
		}

		// Simple admin check - in real app, check roles from database
		if email, ok := userEmail.(string); ok {
			if strings.HasSuffix(email, "@admin.com") {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Forbidden",
			"message": "Admin access required",
		})
		c.Abort()
	}
}
