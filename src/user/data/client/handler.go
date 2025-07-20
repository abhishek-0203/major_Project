package client

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Client represents the structure of a client's profile data from the database.
type Client struct {
	UserID        int    `json:"user_id"`
	Name          string `json:"name"`
	Company       string `json:"company"`
	ContactNumber string `json:"contact_number"`
	Address       string `json:"address"`
	Email         string `json:"email"` // Joined from the 'users' table
	CreatedAt     string `json:"created_at"`
}

// CreateClientPayload is the structure for the request body when creating a new client.
// In a real app, this would be part of a signup flow.
type CreateClientPayload struct {
	Name          string `json:"name" binding:"required"`
	Company       string `json:"company"`
	ContactNumber string `json:"contact_number"`
	Address       string `json:"address"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required,min=8"`
}

// UpdateClientPayload is the structure for the request body when updating a client.
type UpdateClientPayload struct {
	Name          string `json:"name" binding:"required"`
	Company       string `json:"company"`
	ContactNumber string `json:"contact_number"`
	Address       string `json:"address"`
}

// Handler holds the database connection dependency.
type Handler struct {
	DB *sql.DB
}

// NewHandler creates a new handler with the given database connection.
func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

// GetClients retrieves a list of all clients from the database.
func (h *Handler) GetClients(c *gin.Context) {
	query := `
		SELECT c.user_id, c.name, c.company, c.contact_number, c.address, u.email, u.created_at
		FROM clients c
		JOIN users u ON c.user_id = u.id
		ORDER BY c.name;
	`
	rows, err := h.DB.Query(query)
	if err != nil {
		log.Printf("Error querying clients: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve clients"})
		return
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var client Client
		if err := rows.Scan(&client.UserID, &client.Name, &client.Company, &client.ContactNumber, &client.Address, &client.Email, &client.CreatedAt); err != nil {
			log.Printf("Error scanning client row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process client data"})
			return
		}
		clients = append(clients, client)
	}

	c.JSON(http.StatusOK, clients)
}

// CreateClient creates a new user and a corresponding client profile.
// NOTE: This is a complex operation that should be handled within a database transaction.
func (h *Handler) CreateClient(c *gin.Context) {
	var payload CreateClientPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	// Use a transaction to ensure both tables are updated or neither is.
	tx, err := h.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}

	// 1. Create the user record
	userQuery := "INSERT INTO users (email, password_hash, role) VALUES (?, ?, 'client')"
	res, err := tx.Exec(userQuery, payload.Email, string(hashedPassword))
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}
	newUserID, _ := res.LastInsertId()

	// 2. Create the client profile record
	clientQuery := "INSERT INTO clients (user_id, name, company, contact_number, address) VALUES (?, ?, ?, ?, ?)"
	_, err = tx.Exec(clientQuery, newUserID, payload.Name, payload.Company, payload.ContactNumber, payload.Address)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client profile"})
		return
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Client created successfully", "user_id": newUserID})
}

// UpdateClient updates an existing client's profile information.
func (h *Handler) UpdateClient(c *gin.Context) {
	userID := c.Param("id") // Use the user ID from the URL path

	var payload UpdateClientPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE clients SET name = ?, company = ?, contact_number = ?, address = ? WHERE user_id = ?"
	res, err := h.DB.Exec(query, payload.Name, payload.Company, payload.ContactNumber, payload.Address, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update client profile"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client updated successfully"})
}

// DeleteClient removes a user and their associated client profile.
func (h *Handler) DeleteClient(c *gin.Context) {
	userID := c.Param("id")

	// The 'ON DELETE CASCADE' in your schema will automatically delete the client profile.
	query := "DELETE FROM users WHERE id = ? AND role = 'client'"
	res, err := h.DB.Exec(query, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}
