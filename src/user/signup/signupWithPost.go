package signup

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// Helper function to check if user exists
func userExists(email, role string) bool {
	if role == userLocalDb.CLIENT {
		_, exists := userLocalDb.ClientValidUsers[email]
		return exists
	} else if role == userLocalDb.DEVELOPER {
		_, exists := userLocalDb.DeveloperValidUsers[email]
		return exists
	}
	return false
}

func SignUpRequestWithPost(c *gin.Context) {
	var req SignUpRequest

	// Bind the JSON request body to the SignUpRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate role
	if req.Role != userLocalDb.DEVELOPER && req.Role != userLocalDb.CLIENT {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role. Must be 'developer' or 'client'."})
		return
	}

	// Check if user already exists
	if userExists(req.Email, req.Role) {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email and role."})
		return
	}

	// Hash the password securely
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to securely hash password."})
		return
	}

	// Store user data based on role
	if req.Role == userLocalDb.CLIENT {
		userLocalDb.ClientValidUsers[req.Email] = string(hashedPassword)
		userLocalDb.Clients[req.Email] = userLocalDb.Client{
			Email:      req.Email,
			Password:   string(hashedPassword),
			IsVerified: false,
			// Add other fields as needed from req
		}
	} else {
		userLocalDb.DeveloperValidUsers[req.Email] = string(hashedPassword)
		userLocalDb.Developers[req.Email] = userLocalDb.Developer{
			Email:      req.Email,
			Password:   string(hashedPassword),
			IsVerified: false,
			// Add other fields as needed from req
		}
	}

	// Generate email verification token
	verificationToken := uuid.New().String()
	const verificationExpiryMinutes = 30
	verificationExpiry := time.Now().Add(verificationExpiryMinutes * time.Minute)
	userLocalDb.EmailVerificationTokens[verificationToken] = req.Email
	userLocalDb.EmailVerificationExpiry[verificationToken] = verificationExpiry

	// Mock sending email: return verification link in response
	verificationLink := "http://localhost:8080/verify-email?token=" + verificationToken

	c.JSON(http.StatusCreated, gin.H{
		"message":           "User signed up successfully. Please verify your email.",
		"email":             req.Email,
		"role":              req.Role,
		"verification_link": verificationLink,
	})
}
