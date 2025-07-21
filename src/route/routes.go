package route

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/projects"
	"majorProject/src/user/data/developer"
)

func RegisterProjectRoutes(route *gin.RouterGroup) {
	projects.RegisterProject(route)
}

func RegisterDeveloperRoutes(route *gin.RouterGroup) {
	developer.RegisterDeveloper(route)
}
