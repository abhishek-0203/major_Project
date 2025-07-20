package route

import (
	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(route *gin.RouterGroup) {
	RegisterProject(route)
}

func RegisterClientRoutes(route *gin.RouterGroup) {
	RegisterClient(route)
}

func RegisterDeveloperRoutes(route *gin.RouterGroup) {
	RegisterDeveloper(route)
}
