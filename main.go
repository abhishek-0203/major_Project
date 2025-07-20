package main

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/route"
	"majorProject/src/user/forgot"
	"majorProject/src/user/login"
	"majorProject/src/user/middleware"
	"majorProject/src/user/signup"
	"majorProject/src/user/verify"
	"majorProject/src/video"
)

func main() {
	router := gin.Default()

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

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
