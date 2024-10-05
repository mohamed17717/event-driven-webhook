package apis

import (
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create, Read, Update, Delete for Action model

func CreateAction(c *gin.Context) {
	var action models.Action

	if err := c.ShouldBindJSON(&action); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&action).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create action"})
		return
	}

	c.JSON(http.StatusOK, action)
}

func GetActions(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	fmt.Printf("UserId = %s\n\n", user)

	var actions []models.Action

	if err := config.DB.Find(&actions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch actions"})
		return
	}

	c.JSON(http.StatusOK, actions)
}

func GetAction(c *gin.Context) {
	var action models.Action
	id := c.Param("id")

	if err := config.DB.First(&action, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Action not found"})
		return
	}

	c.JSON(http.StatusOK, action)
}

func UpdateAction(c *gin.Context) {
	var action models.Action
	id := c.Param("id")

	if err := config.DB.First(&action, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Action not found"})
		return
	}

	if err := c.ShouldBindJSON(&action); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&action)
	c.JSON(http.StatusOK, action)
}

func DeleteAction(c *gin.Context) {
	var action models.Action
	id := c.Param("id")

	if err := config.DB.First(&action, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Action not found"})
		return
	}

	config.DB.Delete(&action)
	c.JSON(http.StatusOK, gin.H{"message": "Action deleted"})
}

func ActionRoutes() {
	protected := ProtectedRoute()

	protected.POST("/actions", CreateAction)
	protected.GET("/actions", GetActions)
	protected.GET("/actions/:id", GetAction)
	protected.PUT("/actions/:id", UpdateAction)
	protected.DELETE("/actions/:id", DeleteAction)
}
