package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TestController handles test endpoints for development
type TestController struct{}

// NewTestController creates a new test controller
func NewTestController() *TestController {
	return &TestController{}
}

// TestPanic creates an intentional panic for testing panic recovery
func (tc *TestController) TestPanic(c *gin.Context) {
	// Get panic type from query parameter
	panicType := c.DefaultQuery("type", "string")

	switch panicType {
	case "nil":
		// Nil pointer dereference
		var p *string
		fmt.Println(*p) // This will panic
	case "slice":
		// Index out of bounds
		arr := []string{"a", "b", "c"}
		fmt.Println(arr[10]) // This will panic
	case "map":
		// Map access panic
		var m map[string]string
		m["key"] = "value" // This will panic (nil map)
	case "divide":
		// Division by zero (though Go handles this differently)
		x := 1
		y := 0
		if y == 0 {
			panic("division by zero")
		}
		result := x / y
		fmt.Println(result)
	case "custom":
		// Custom panic message
		message := c.DefaultQuery("message", "This is a test panic!")
		panic(message)
	default:
		// Default string panic
		panic("This is a test panic to verify panic recovery!")
	}

	// This line should never be reached
	c.JSON(http.StatusOK, gin.H{
		"message": "This should not be returned if panic occurs",
	})
}

// TestError creates a regular error (not panic) for comparison
func (tc *TestController) TestError(c *gin.Context) {
	// Add error to gin context
	c.Error(fmt.Errorf("this is a test error, not a panic"))
	
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   "Bad Request",
		"message": "This is a controlled error for testing",
	})
}

// TestSuccess returns a successful response
func (tc *TestController) TestSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Test endpoint working successfully",
		"status":  "ok",
	})
}
