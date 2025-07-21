package route

import (
	"database/sql"
	"majorProject/src/user/data/client" // Import the client package

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes sets up the routes for the client API endpoints.
// It accepts the Gin engine and the database connection.
func RegisterClientRoutes(router *gin.Engine, db *sql.DB) {
	// Create a new handler, injecting the database connection.
	h := client.NewHandler(db)

	// Group routes under the "/clients" path.
	group := router.Group("/clients")
	{
		group.POST("/", h.CreateClient)
		group.GET("/", h.GetClients)
		// It's better to use a unique ID in the path for updates and deletions.
		group.PUT("/:id", h.UpdateClient)
		group.DELETE("/:id", h.DeleteClient)
	}
}
