package handler

import (
	"fmt"
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddEvent(c *gin.Context) {
	eventRegisterRequest := models.RegisterEvent{}

	err := c.ShouldBindJSON(&eventRegisterRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	newEvent, err := models.RecordNewEvent(&eventRegisterRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": newEvent, "success": true})
}

func (h *Handler) ShowEvent(c *gin.Context) {

	userId := c.GetUint("user")
	user := models.GetUser(userId)
	fmt.Println(user)

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
	eventUpdateData := models.UpdateEvent{}

	err := c.ShouldBindJSON(&eventUpdateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	eventId, err := h.getPathParamUint(c, "eventId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	updatedEvent, err := models.UpdateEventRecord(&eventUpdateData, eventId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": updatedEvent, "success": true})
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

func (h *Handler) GetEventsInLocation(c *gin.Context) {
	locationId, err := h.getPathParamUint(c, "locationId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	eventsInLocation, err := models.FindAllEventsInLocation(*locationId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": eventsInLocation, "success": true})
}

func (h *Handler) GetEventsInArea(c *gin.Context) {

	lat1 := c.Query("lat1")
	lng1 := c.Query("lng1")

	lat2 := c.Query("lat2")
	lng2 := c.Query("lng2")

	eventsInArea, err := models.FindEventsInArea(lat1, lng1, lat2, lng2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event": eventsInArea, "success": true})
}
