package route

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/user/data/client"
)

// RegisterClientRoutes sets up the client routes using the client package's RegisterClient function.
func RegisterClientRoutes(route *gin.RouterGroup) {
	client.RegisterClient(route)
}
