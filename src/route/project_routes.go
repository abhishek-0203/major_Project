package route

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"majorProject/src/projects"

	"github.com/gin-gonic/gin"
)

// Utility function to write back to file
func saveProjectsToFile(filePath string, projectsList []projects.Project) error {
	data, err := json.MarshalIndent(projectsList, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func RegisterProject(route *gin.RouterGroup) {
	filePath := "Doc/project/project.json"

	// GET all projects
	route.GET("/projects", func(c *gin.Context) {
		projectsList, err := projects.LoadProjectsFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"projects": projectsList})
	})

	// POST: Create a new project
	route.POST("/createProject", func(c *gin.Context) {
		var newProject projects.Project
		if err := c.ShouldBindJSON(&newProject); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project data: " + err.Error()})
			return
		}

		// Validate required fields
		if newProject.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Project title is required"})
			return
		}
		if newProject.Description == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Project description is required"})
			return
		}
		if newProject.Budget <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Project budget must be greater than 0"})
			return
		}

		// Set default values
		if newProject.Status == "" {
			newProject.Status = "open"
		}

		// Get current timestamp
		currentTime := time.Now().Format(time.RFC3339)
		newProject.CreatedAt = currentTime

		// Generate a unique project ID
		newProject.ProjectID = "proj_" + strconv.FormatInt(time.Now().UnixNano(), 10)

		projectsList, err := projects.LoadProjectsFromFile(filePath)
		if err != nil {
			// If file doesn't exist, create new slice
			if os.IsNotExist(err) {
				projectsList = []projects.Project{}
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load projects: " + err.Error()})
				return
			}
		}

		projectsList = append(projectsList, newProject)
		if err := saveProjectsToFile(filePath, projectsList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Project added successfully!"})
	})

	// PUT: Update project by index
	route.PUT("/updateProject/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
			return
		}

		var updatedProject projects.Project
		if err := c.ShouldBindJSON(&updatedProject); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		projectsList, err := projects.LoadProjectsFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if index < 0 || index >= len(projectsList) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}

		projectsList[index] = updatedProject
		if err := saveProjectsToFile(filePath, projectsList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
	})

	// DELETE: Remove project by index
	route.DELETE("/deleteProject/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
			return
		}

		projectsList, err := projects.LoadProjectsFromFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if index < 0 || index >= len(projectsList) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}

		projectsList = append(projectsList[:index], projectsList[index+1:]...)
		if err := saveProjectsToFile(filePath, projectsList); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
	})
}
