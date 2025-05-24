package signup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"majorProject/src/user/userLocalDb"
)

// SignUpRequest defines the structure for the JSON payload for signup
type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`    // Added binding tags for validation
	Password string `json:"password" binding:"required,min=6"` // Example: min password length
	//Phone    string `json:"phone" binding:"required"`
	Role string `json:"role" binding:"required"`
}

func SignUpRequestWithPost(c *gin.Context) {
	var req SignUpRequest

	// Bind the JSON request body to the SignUpRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return specific error message from binding
		return
	}

	// Validating the role
	if req.Role != userLocalDb.DEVELOPER && req.Role != userLocalDb.CLIENT {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role. Must be 'developer' or 'client'."})
		return
	}

	// Checking if the email already exists or not
	if req.Role == userLocalDb.CLIENT {
		if _, exists := userLocalDb.ClientValidUsers[req.Email]; exists {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email"})
			return
		}
	} else { // Assuming it's a developer if not a client (due to previous validation)
		if _, exists := userLocalDb.DeveloperValidUsers[req.Email]; exists {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email"})
			return
		}
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}

	// Storing user data based on role
	if req.Role == userLocalDb.CLIENT {
		userLocalDb.ClientValidUsers[req.Email] = string(hashedPassword)
	} else {
		userLocalDb.DeveloperValidUsers[req.Email] = string(hashedPassword)
	}
	//	userLocalDb.RegisteredUser[req.Email] = req.Phone // Assuming RegisteredUser stores email -> phone for all types

	c.JSON(http.StatusCreated, gin.H{
		"message": "User signed up successfully",
		"email":   req.Email,
		"role":    req.Role,
	})
}
