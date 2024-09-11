package handlers

import (
	"github.com/gin-gonic/gin"
	"marcopolo/internal/config"
	"marcopolo/pkg/models"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	userDB := config.DB

	// Get all records
	if err := userDB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
