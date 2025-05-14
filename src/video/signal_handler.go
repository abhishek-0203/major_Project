package video

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JoinCallHandler(c *gin.Context) {
	var req JoinCallRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	//if !ValidateToken(req.Token) {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
	//	return
	//}

	resp := JoinCallResponse{
		Message: "Joined call successfully",
		RoomID:  req.RoomID,
		UserID:  req.UserID,
	}

	c.JSON(http.StatusOK, resp)
}
