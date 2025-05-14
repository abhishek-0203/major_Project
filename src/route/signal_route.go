package route

import (
	"majorProject/src/video"

	"github.com/gin-gonic/gin"
)

func RegisterVideoRoutes(router *gin.Engine) {
	router.POST("/join-call", video.JoinCallHandler)
}
