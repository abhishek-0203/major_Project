package main

import (
	"github.com/gin-gonic/gin"
	routes "majorProject/src/route"
	"majorProject/src/user/forgot"
	"majorProject/src/user/login"
	"majorProject/src/user/signup"
	"majorProject/src/video"
)

func main() {
	route := gin.Default()

	route.POST("/login", login.LoginRequestWithPost)
	route.GET("/signup", signup.SignUpRequestWithGet)
	route.GET("/forgotpassword", forgot.ForgotPasswordWithGet)
	route.GET("/reset-password", forgot.ResetPasswordWithGet)

	routes.RegisterProject(route)
	routes.RegisterClient(route)
	routes.RegisterDeveloper(route)

	routes.RegisterChat(route)
	video.RegisterVideoCallRoutes(route)
	routes.RegisterVideoRoutes(route)
	routes.RegisterPaymentRoutes(route)
	/*
		routes.RegisterProjectRoutes(route)
		routes.RegisterClientRoutes(route)
		routes.RegisterDeveloperRoutes(route)
	*/
	route.Run() // listen and serve on 0.0.0.0:8080
}
