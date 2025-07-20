package forgot

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"majorProject/src/user/userLocalDb"
	"net/http"
	"net/url"
	"time"
)

func ResetPasswordWithGet(c *gin.Context) {
	query := c.Request.URL.RawQuery
	params, _ := url.ParseQuery(query)

	token := params.Get("token")
	newPassword := params.Get("newPassword")
	userType := params.Get("userType")

	if token == "" || newPassword == "" || userType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token, new password, and user type are required"})
		return
	}

	// Validate user type
	if userType != "client" && userType != "developer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user type. Must be 'client' or 'developer'."})
		return
	}

	// Check token existence and expiry
	email, tokenExists := userLocalDb.ResetTokens[token]
	expiry, timeExists := userLocalDb.TokenExpiry[token]

	if !tokenExists || !timeExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired token. Please request a new password reset."})
		return
	}

	if time.Now().After(expiry) {
		delete(userLocalDb.ResetTokens, token)
		delete(userLocalDb.TokenExpiry, token)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token has expired. Please request a new password reset."})
		return
	}

	// Hash new password securely
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to securely hash password."})
		return
	}

	// Update password based on user type
	if userType == "client" {
		userLocalDb.ClientValidUsers[email] = string(hashedPassword)
	} else {
		userLocalDb.DeveloperValidUsers[email] = string(hashedPassword)
	}

	// Invalidate token
	delete(userLocalDb.ResetTokens, token)
	delete(userLocalDb.TokenExpiry, token)

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successful"})
}
