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
	event, err := models.GetEvent(*eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

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

	event, err := models.GetEvent(*eventId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	event.UpdateEventFields(&eventUpdate)
	updatedEvent, err := models.UpdateEventRecord(event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": updatedEvent, "success": true})
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

	err = models.DeleteEvent(*eventId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *Handler) GetAllEvents(c *gin.Context) {
	allEvents, err := models.FindAllEvents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event": allEvents, "success": true})
}
