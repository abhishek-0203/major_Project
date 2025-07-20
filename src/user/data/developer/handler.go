package developer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterDeveloper(route *gin.RouterGroup) {
	group := route.Group("")
	group.GET("/developers", GetDevelopers)
	group.POST("/developers", CreateDeveloper)
	group.PUT("/developers/:email", UpdateDeveloper)
	group.DELETE("/developers/:email", DeleteDeveloper)
}

func GetDevelopers(c *gin.Context) {
	developers, err := LoadDevelopers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, developers)
}

func CreateDeveloper(c *gin.Context) {
	var newDeveloper Developer
	if err := c.BindJSON(&newDeveloper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	developers, _ := LoadDevelopers()
	developers = append(developers, newDeveloper)
	SaveDevelopers(developers)
	c.JSON(http.StatusOK, gin.H{"message": "Developer created successfully"})
}

func UpdateDeveloper(c *gin.Context) {
	email := c.Param("email")
	var updatedDeveloper Developer
	if err := c.BindJSON(&updatedDeveloper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	developers, _ := LoadDevelopers()
	for i, dev := range developers {
		if dev.Email == email {
			developers[i] = updatedDeveloper
			SaveDevelopers(developers)
			c.JSON(http.StatusOK, gin.H{"message": "Developer updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Developer not found"})
}

func DeleteDeveloper(c *gin.Context) {
	email := c.Param("email")
	developers, _ := LoadDevelopers()
	for i, dev := range developers {
		if dev.Email == email {
			developers = append(developers[:i], developers[i+1:]...)
			SaveDevelopers(developers)
			c.JSON(http.StatusOK, gin.H{"message": "Developer deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Developer not found"})
}
