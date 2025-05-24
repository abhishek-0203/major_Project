package login

import (
	"majorProject/src/user/data/developer"
	"net/http"

	"github.com/gin-gonic/gin"
	"majorProject/src/user/data/client" // Import the client package
	"majorProject/src/user/userLocalDb"
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

	var passwd string
	var isExist bool

	if req.Role == userLocalDb.DEVELOPER {
		passwd, isExist = userLocalDb.DeveloperValidUsers[req.Email]
	} else if req.Role == userLocalDb.CLIENT {
		passwd, isExist = userLocalDb.ClientValidUsers[req.Email]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	if isExist {
		if passwd == req.Password { // This is the block that should contain the success logic
			loginUser = userLocalDb.LoginUser{
				Email:    req.Email,
				Password: req.Password,
				Phone:    req.Phone,
				Status:   true,
				Message:  "Logged in successfully.",
				Token:    req.Email,
				Role:     req.Role,
			}

			// --- Fetching and adding profile to response ---
			responsePayload := gin.H{
				"data": loginUser,
			}

			if req.Role == userLocalDb.CLIENT {
				filePath := "Doc/project/client.json"
				clientsList, err := client.LoadClientsFromFile(filePath)
				if err != nil {
					responsePayload["profile_error"] = "Failed to load client profile: " + err.Error()
				} else {
					var foundClient *client.Client
					for i := range clientsList {
						if clientsList[i].Email == req.Email { // Assuming Client struct has an Email field
							foundClient = &clientsList[i]
							break
						}
					}
					if foundClient != nil {
						responsePayload["profile"] = *foundClient // Add the client profile
					} else {
						responsePayload["profile_message"] = "Client profile not found for this email."
					}
				}
			} else if req.Role == userLocalDb.DEVELOPER {
				filePath := "Doc/project/developer.json"
				developersList, err := developer.LoadDevelopersFromFile(filePath)
				if err != nil {
					responsePayload["profile_error"] = "Failed to load developer profile: " + err.Error()
				} else {
					var foundDeveloper *developer.Developer
					for i := range developersList {
						if developersList[i].Email == req.Email {
							foundDeveloper = &developersList[i]
							break
						}
					}
					if foundDeveloper != nil {
						responsePayload["profile"] = *foundDeveloper
					} else {
						responsePayload["profile_message"] = "Developer profile not found for this email."
					}
				}
			}
			// --- End of profile fetching ---

			// These lines were incorrectly placed outside the 'if' block
			userLocalDb.LoggedInUserList = append(userLocalDb.LoggedInUserList, loginUser)
			c.JSON(http.StatusOK, responsePayload) // Send the enriched response
		} else { // Password is incorrect
			loginUser = userLocalDb.LoginUser{Email: req.Email, Status: false, Message: "Password is incorrect."}
			c.JSON(http.StatusOK, gin.H{"data": loginUser})
		}
	} else { // Login is not authorized (email not found)
		loginUser = userLocalDb.LoginUser{
			Email:   req.Email,
			Status:  false,
			Message: "Login is not authorized for " + req.Role + " role.",
		}
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"data": loginUser})
	}
}
