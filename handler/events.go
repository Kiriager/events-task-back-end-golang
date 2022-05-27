package handler

import (
	"fmt"
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

func (h *Handler) ShowEvent(c *gin.Context) {

	eventId, err := h.getPathParamUint(c, "eventId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	event := models.GetEvent(*eventId)

	c.JSON(http.StatusOK, gin.H{"event": event, "success": true})
}

func (h *Handler) UpdateEvent(c *gin.Context) {
	eventUpdate := models.UpdateEvent{}
	err := c.ShouldBindJSON(&eventUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	eventId, err := h.getPathParamUint(c, "eventId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	event := models.GetEvent(*eventId)
	event.UpdateEventFields(&eventUpdate)
	event.UpdateEventRecord()

	c.JSON(http.StatusOK, gin.H{"event": event, "success": true})
}

func (h *Handler) test(c *gin.Context) {
	fmt.Println("Hello from test")
}

func (h *Handler) DeleteEvent(c *gin.Context) {

	eventId, err := h.getPathParamUint(c, "eventId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	models.DeleteEvent(*eventId)
	event := models.GetEvent(*eventId)

	c.JSON(http.StatusOK, gin.H{"event": event, "success": true})
}
