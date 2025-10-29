package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterProject(route *gin.RouterGroup) {
	group := route.Group("/projects")
	group.GET("/getProjects", GetProjects)
	group.POST("/createProject", CreateProject)
	group.PUT("/updateProject/:title", UpdateProject)
	group.DELETE("/deleteProject/:title", DeleteProject)
}

func GetProjects(c *gin.Context) {
	projects, err := LoadProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func CreateProject(c *gin.Context) {
	var newProject Project
	if err := c.BindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var err error
	projects, _ := LoadProjects()
	projects = append(projects, newProject)
	err = SaveProjects(projects)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project created successfully"})
}

func UpdateProject(c *gin.Context) {
	projectID := c.Param("id")
	var updatedProject Project
	if err := c.BindJSON(&updatedProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	projects, _ := LoadProjects()
	for i, proj := range projects {
		if proj.ProjectID == projectID {
			// Preserve the original ID and creation time
			updatedProject.ProjectID = proj.ProjectID
			updatedProject.CreatedAt = proj.CreatedAt
			projects[i] = updatedProject
			err := SaveProjects(projects)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Project updated", "project": updatedProject})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

func DeleteProject(c *gin.Context) {
	projectID := c.Param("id")
	projects, _ := LoadProjects()
	for i, proj := range projects {
		if proj.ProjectID == projectID {
			projects = append(projects[:i], projects[i+1:]...)
			err := SaveProjects(projects)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}
