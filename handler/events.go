package handler

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddEvent(c *gin.Context) {
	eventCreate := models.CreateEvent{}
	err := c.ShouldBindJSON(&eventCreate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	event, err := models.AddEvent(&eventCreate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event, "success": true})
}
