package login

import (
	"majorProject/src/user/auth"
	"majorProject/src/user/data/client" // Import the client package
	"majorProject/src/user/data/developer"
	"majorProject/src/user/userLocalDb"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	// Validate role
	if req.Role != userLocalDb.DEVELOPER && req.Role != userLocalDb.CLIENT {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role. Please specify 'developer' or 'client'."})
		return
	}
	userLocalDb.CURRENT_USER_ROLE = req.Role

	var passwd string
	var isExist bool
	if req.Role == userLocalDb.DEVELOPER {
		passwd, isExist = userLocalDb.DeveloperValidUsers[req.Email]
	} else {
		passwd, isExist = userLocalDb.ClientValidUsers[req.Email]
	}

	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not exist or wrong role."})
		return
	}

	// Use bcrypt to compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(passwd), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password."})
		return
	}

	// Check if user is verified
	isVerified := false
	if req.Role == userLocalDb.CLIENT {
		if cStruct, ok := userLocalDb.Clients[req.Email]; ok {
			isVerified = cStruct.IsVerified
		}
	} else if req.Role == userLocalDb.DEVELOPER {
		if dStruct, ok := userLocalDb.Developers[req.Email]; ok {
			isVerified = dStruct.IsVerified
		}
	}

	if !isVerified {
		c.JSON(http.StatusForbidden, gin.H{"error": "Email not verified. Please verify your email before logging in."})
		return
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(req.Email, req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token."})
		return
	}

	loginUser := userLocalDb.LoginUser{
		Email:    req.Email,
		Password: "", // Do not return password
		Phone:    req.Phone,
		Status:   true,
		Message:  "Logged in successfully.",
		Token:    token, // Use JWT token here
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
}
