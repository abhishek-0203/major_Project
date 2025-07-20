package route

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/chat"
)

func RegisterChat(route *gin.RouterGroup) {
	route.GET("/ws/chat", func(c *gin.Context) {
		chat.ChatHandler(c.Writer, c.Request)
	})
}
