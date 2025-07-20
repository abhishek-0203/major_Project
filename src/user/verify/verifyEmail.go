package verify

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"majorProject/src/user/userLocalDb"
)

func VerifyEmailWithGet(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Verification token is required."})
		return
	}

	email, tokenExists := userLocalDb.EmailVerificationTokens[token]
	expiry, expiryExists := userLocalDb.EmailVerificationExpiry[token]

	if !tokenExists || !expiryExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired verification token."})
		return
	}

	if time.Now().After(expiry) {
		delete(userLocalDb.EmailVerificationTokens, token)
		delete(userLocalDb.EmailVerificationExpiry, token)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Verification token has expired."})
		return
	}

	// Mark user as verified
	if client, exists := userLocalDb.Clients[email]; exists {
		client.IsVerified = true
		userLocalDb.Clients[email] = client
	} else if developer, exists := userLocalDb.Developers[email]; exists {
		developer.IsVerified = true
		userLocalDb.Developers[email] = developer
	}

	// Remove token after successful verification
	delete(userLocalDb.EmailVerificationTokens, token)
	delete(userLocalDb.EmailVerificationExpiry, token)

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully."})
}
