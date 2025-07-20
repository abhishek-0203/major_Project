package main

import (
	"database/sql"
	"log"
	"majorProject/database" // Import the database package
	routes "majorProject/src/route"
	"majorProject/src/user/forgot"
	"majorProject/src/user/login"
	"majorProject/src/user/signup"
	"majorProject/src/video"

	"github.com/gin-gonic/gin"
)

func main() {
	// --- Database Initialization ---
	// This part is correct.
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// --- Gin Router Setup ---
	router := gin.Default()

	// --- Route Registration with Dependency Injection ---
	// Now, we pass the 'db' object to the functions that set up routes.
	// This makes the database connection available to the handlers.

	// For individual handlers, you need to adapt them to accept the db connection.
	// The standard way is to use a closure.
	router.POST("/login", login.LoginHandler(db))
	router.POST("/signup", signup.SignUpHandler(db))

	// These handlers seem to be for rendering pages, so they might not need the DB.
	// If they do, apply the same pattern.
	router.GET("/forgotpassword", forgot.ForgotPasswordWithGet)
	router.GET("/reset-password", forgot.ResetPasswordWithGet)

	// You must update your route registration functions to accept the 'db' object.
	routes.RegisterProjectRoutes(router, db)
	routes.RegisterClientRoutes(router, db)
	routes.RegisterDeveloperRoutes(router, db)
	routes.RegisterChatRoutes(router, db)
	video.RegisterVideoCallRoutes(router) // This might not need the DB
	routes.RegisterVideoRoutes(router, db)
	routes.RegisterPaymentRoutes(router, db)

	// --- Start Server ---
	log.Println("Starting server on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
