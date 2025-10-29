package main

import (
	"majorProject/src/route"
	"majorProject/src/user/forgot"
	"majorProject/src/user/login"
	"majorProject/src/user/middleware"
	"majorProject/src/user/signup"
	"majorProject/src/user/verify"
	"majorProject/src/video"

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
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Public routes (no JWT required)
	router.POST("/login", login.LoginRequestWithPost)
	router.POST("/signup", signup.SignUpRequestWithPost)
	router.GET("/forgotpassword", forgot.ForgotPasswordWithGet)
	router.GET("/reset-password", forgot.ResetPasswordWithGet)
	router.GET("/verify-email", verify.VerifyEmailWithGet)

	// Protected routes (JWT required)
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		route.RegisterProject(protected)
		route.RegisterClient(protected)
		route.RegisterDeveloper(protected)
		route.RegisterChat(protected)
		video.RegisterVideoCallRoutes(protected)
		route.RegisterPaymentRoutes(protected)

		// Add other protected routes here
	}

	err := router.Run(":8081")
	if err != nil {
		panic(err)
	}
}
