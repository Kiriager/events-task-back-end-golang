package handler

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddLocation(c *gin.Context) {
	locationData := models.RegisterLocation{}
	err := c.ShouldBindJSON(&locationData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	newLocation, err := models.RecordLocation(&locationData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"location": newLocation, "success": true})

}
