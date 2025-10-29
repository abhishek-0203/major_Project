package main

import (
	"log"
	"majorProject/database"
	"majorProject/src/route"
	"majorProject/src/user/forgot"
	login "majorProject/src/user/login"
	jwtMiddleware "majorProject/src/user/middleware"
	signup "majorProject/src/user/signup"
	"majorProject/src/user/verify"
	"majorProject/src/video"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// Initialize MongoDB connection
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer database.CloseDB()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Public routes (no JWT required)
	router.POST("/api/login", login.LoginRequestWithPost)
	router.POST("/api/signup", signup.SignUpRequestWithPost)
	router.POST("/api/forgotPassword", forgot.ForgotPasswordWithGet)
	router.GET("/api/verify", verify.VerifyEmailWithGet)
	router.POST("/api/resetPassword", forgot.ResetPasswordWithGet)

	// Protected routes (JWT required)
	var protected *gin.RouterGroup
	protected = router.Group("/api")
	protected.Use(jwtMiddleware.JWTAuthMiddleware())

	// Register routes for projects, clients, developers, etc.
	route.RegisterProject(protected)
	route.RegisterClient(protected)
	route.RegisterDeveloper(protected)
	route.RegisterChat(protected)
	video.RegisterVideoCallRoutes(protected)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
	})

	// Start the server
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
