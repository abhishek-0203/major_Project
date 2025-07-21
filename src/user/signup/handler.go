package signup

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// SignUpPayload defines the structure for the signup request body.
// It includes validation tags to ensure data quality.
type SignUpPayload struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required,oneof=client developer"`
}

// SignUpHandler creates a gin.HandlerFunc that handles user registration.
// It takes the database connection as a dependency.
func SignUpHandler(db *sql.DB) gin.HandlerFunc {
	// The returned function is a "closure" that has access to the 'db' variable.
	return func(c *gin.Context) {
		var payload SignUpPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			// Gin's binding validation provides descriptive errors.
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// --- Securely Hash the Password ---
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Error hashing password: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process registration"})
			return
		}

		// --- Use a Database Transaction ---
		// A transaction ensures that we either create both the user and their profile,
		// or we create neither. This prevents inconsistent data.
		tx, err := db.Begin()
		if err != nil {
			log.Printf("Error beginning transaction: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process registration"})
			return
		}
		// Defer a rollback in case of a panic. The explicit commit at the end will override this.
		defer tx.Rollback()

		// 1. Insert into the main 'users' table
		userQuery := "INSERT INTO users (email, password_hash, role) VALUES (?, ?, ?)"
		res, err := tx.Exec(userQuery, payload.Email, string(hashedPassword), payload.Role)
		if err != nil {
			// This error is likely due to a duplicate email (UNIQUE constraint).
			c.JSON(http.StatusConflict, gin.H{"error": "A user with this email already exists"})
			return
		}
		newUserID, _ := res.LastInsertId()

		// 2. Insert into the role-specific profile table ('clients' or 'developers')
		if payload.Role == "client" {
			profileQuery := "INSERT INTO clients (user_id, name) VALUES (?, ?)"
			if _, err := tx.Exec(profileQuery, newUserID, payload.Name); err != nil {
				log.Printf("Error creating client profile: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client profile"})
				return
			}
		} else if payload.Role == "developer" {
			profileQuery := "INSERT INTO developers (user_id, name) VALUES (?, ?)"
			if _, err := tx.Exec(profileQuery, newUserID, payload.Name); err != nil {
				log.Printf("Error creating developer profile: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create developer profile"})
				return
			}
		}

		// --- Commit the Transaction ---
		// If all steps were successful, commit the changes to the database.
		if err := tx.Commit(); err != nil {
			log.Printf("Error committing transaction: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to finalize registration"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User registered successfully",
			"user_id": newUserID,
			"role":    payload.Role,
		})
	}
}
