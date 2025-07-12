package controllers

import (
	"net/http"
	"strconv"

	"golang-domain-driven-design/internal/application/services"
	"golang-domain-driven-design/internal/application/usecases"
	apperrors "golang-domain-driven-design/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserController handles HTTP requests for user operations
type UserController struct {
	userUseCase usecases.UserUseCase
	jwtService  services.JWTService
}

// NewUserController creates a new user controller
func NewUserController(userUseCase usecases.UserUseCase, jwtService services.JWTService) *UserController {
	return &UserController{
		userUseCase: userUseCase,
		jwtService:  jwtService,
	}
}

// CreateUser handles POST /api/v1/users
func (c *UserController) CreateUser(ctx *gin.Context) {
	var req usecases.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	user, err := c.userUseCase.CreateUser(ctx.Request.Context(), req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    user,
	})
}

// RegisterUser handles user registration
func (uc *UserController) RegisterUser(c *gin.Context) {
	var req usecases.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userUseCase.CreateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token for the new user
	authResponse, err := uc.jwtService.GenerateToken(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(201, authResponse)
}

// GetUser handles GET /api/v1/users/:id
func (c *UserController) GetUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid user ID format",
		})
		return
	}

	user, err := c.userUseCase.GetUser(ctx.Request.Context(), id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User retrieved successfully",
		"data":    user,
	})
}

// UpdateUser handles PUT /api/v1/users/:id
func (c *UserController) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid user ID format",
		})
		return
	}

	var req usecases.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	user, err := c.userUseCase.UpdateUser(ctx.Request.Context(), id, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"data":    user,
	})
}

// DeleteUser handles DELETE /api/v1/users/:id
func (c *UserController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid user ID format",
		})
		return
	}

	err = c.userUseCase.DeleteUser(ctx.Request.Context(), id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

// ListUsers handles GET /api/v1/users
func (c *UserController) ListUsers(ctx *gin.Context) {
	offsetParam := ctx.DefaultQuery("offset", "0")
	limitParam := ctx.DefaultQuery("limit", "10")

	offset, err := strconv.Atoi(offsetParam)
	if err != nil || offset < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid offset parameter",
		})
		return
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 || limit > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid limit parameter (must be between 1 and 100)",
		})
		return
	}

	response, err := c.userUseCase.ListUsers(ctx.Request.Context(), offset, limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Users retrieved successfully",
		"data":    response,
	})
}

// AuthenticateUser handles POST /api/v1/auth/login
func (c *UserController) AuthenticateUser(ctx *gin.Context) {
	var req usecases.AuthRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	user, err := c.userUseCase.AuthenticateUser(ctx.Request.Context(), req.Email, req.Password)
	if err != nil {
		handleError(ctx, err)
		return
	}

	// Generate JWT token
	response, err := c.jwtService.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to generate authentication token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Authentication successful",
		"data":    response,
	})
}

// RefreshToken handles POST /api/v1/auth/refresh
func (c *UserController) RefreshToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Authorization header is required",
		})
		return
	}

	tokenString := services.ExtractTokenFromHeader(authHeader)
	if tokenString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid authorization header format",
		})
		return
	}

	response, err := c.jwtService.RefreshToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid or expired token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Token refreshed successfully",
		"data":    response,
	})
}

// GetProfile handles GET /api/v1/auth/profile
func (c *UserController) GetProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "User not authenticated",
		})
		return
	}

	id, ok := userID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Invalid user ID format",
		})
		return
	}

	user, err := c.userUseCase.GetUser(ctx.Request.Context(), id)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Profile retrieved successfully",
		"data":    user,
	})
}

// handleError handles different types of errors and returns appropriate HTTP responses
func handleError(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case *apperrors.ValidationError:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation Error",
			"message": e.Message,
			"field":   e.Field,
		})
	case *apperrors.NotFoundError:
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": e.Message,
		})
	case *apperrors.UnauthorizedError:
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": e.Message,
		})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "An unexpected error occurred",
		})
	}
}
