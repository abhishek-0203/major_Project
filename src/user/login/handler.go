package login

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginHandler creates a gin.HandlerFunc that handles the login request.
// It takes the database connection as a dependency.
func LoginHandler(db *sql.DB) gin.HandlerFunc {
	// The returned function is a "closure", it has access to the 'db' variable.
	return func(c *gin.Context) {
		var payload LoginPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		var storedHash string
		query := "SELECT password_hash FROM users WHERE email = ?"
		err := db.QueryRow(query, payload.Email).Scan(&storedHash)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		// Compare the provided password with the stored hash
		err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(payload.Password))
		if err != nil {
			// Passwords don't match
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Login successful (in a real app, you would generate a JWT token here)
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	}
}
