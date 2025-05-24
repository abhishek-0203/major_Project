package login

import (
	"github.com/gin-gonic/gin"
	"majorProject/src/user/userLocalDb"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"` // New role parameter
}

func LoginRequestWithPost(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if req.Role == userLocalDb.DEVELOPER {
		userLocalDb.CURRENT_USER_ROLE = req.Role
	} else if req.Role == userLocalDb.CLIENT {
		userLocalDb.CURRENT_USER_ROLE = req.Role
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role. Please specify 'developer' or 'client'."})
		return
	}

	var loginUser userLocalDb.LoginUser

	passwd, isExist := userLocalDb.ValidUsers[req.Email]
	if isExist {
		if passwd == req.Password {
			loginUser = userLocalDb.LoginUser{
				Email:    req.Email,
				Password: req.Password,
				Phone:    req.Phone,
				Status:   true,
				Message:  "Logged in successfully.",
				Token:    req.Email,
			}
			userLocalDb.LoggedInUserList = append(userLocalDb.LoggedInUserList, loginUser)
			c.JSON(http.StatusOK, gin.H{"data": loginUser})
		} else {
			loginUser = userLocalDb.LoginUser{Email: req.Email, Status: false, Message: "Password is incorrect."}
			c.JSON(http.StatusOK, gin.H{"data": loginUser})
		}
	} else {
		loginUser = userLocalDb.LoginUser{Email: req.Email, Status: false, Message: "Login is not authorized."}
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"data": loginUser})
	}
}
