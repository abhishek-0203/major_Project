package forgot

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"majorProject/src/user/userLocalDb"
	"net/http"
	"regexp"
	"time"
)

func ForgotPasswordWithGet(c *gin.Context) {
	// Use c.Query for cleaner query param parsing
	email := c.Query("email")
	userType := c.Query("userType")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	// Validate email format (simple regex)
	if !isValidEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	if userType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User type is required"})
		return
	}

	// Check if user exists
	var exists bool
	if userType == "client" {
		_, exists = userLocalDb.ClientValidUsers[email]
	} else if userType == "developer" {
		_, exists = userLocalDb.DeveloperValidUsers[email]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type"})
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No user found with this email"})
		return
	}

	// Generate reset token
	token := uuid.New().String()
	const tokenExpiryMinutes = 15
	expiry := time.Now().Add(tokenExpiryMinutes * time.Minute)

	// Save token (consider hashing in production)
	userLocalDb.ResetTokens[token] = email
	userLocalDb.TokenExpiry[token] = expiry

	// Mock "send email"
	resetLink := "http://localhost:8080/reset-password?token=" + token

	c.JSON(http.StatusOK, gin.H{
		"message":    "Password reset link generated.",
		"expires_in": "15 minutes",
		"reset_link": resetLink, // mock link
	})
}

// isValidEmail checks basic email format
func isValidEmail(email string) bool {
	// Simplified regex, removed redundant escapes
	return regexp.MustCompile(`^[\w.-]+@[\w.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}
