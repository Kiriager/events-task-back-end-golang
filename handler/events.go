package handler

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddEvent(c *gin.Context) {
	authorizedUserId := c.GetUint("user")
	authorizedUser := models.GetUser(authorizedUserId)

	if authorizedUser.Role != models.Admin && authorizedUser.Role != models.SuperAdmin {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "current user does't have rights to perform action", "success": false})
		return
	}

	eventRegisterRequest := models.RegisterEvent{}

	err := c.ShouldBindJSON(&eventRegisterRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	newEvent, err := models.RecordNewEvent(&eventRegisterRequest, authorizedUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": newEvent, "success": true})
}

func (h *Handler) ShowEvent(c *gin.Context) {

	//userId := c.GetUint("user")
	//user := models.GetUser(userId)
	//fmt.Println(user)

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

	err = models.GetDB().Where("id = ?", eventId).Preload("Location").First(event).Error
	err = models.GetDB().Where("id = ?", eventId).Preload("Users").First(event).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event, "success": true})
}

func (h *Handler) UpdateEvent(c *gin.Context) {
	authorizedUserId := c.GetUint("user")
	authorizedUser := models.GetUser(authorizedUserId)

	if authorizedUser.Role != models.Admin && authorizedUser.Role != models.SuperAdmin {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "current user does't have rights to perform action", "success": false})
		return
	}

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

	err = models.GetDB().Where("id = ?", eventId).Preload("Location").First(updatedEvent).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	err = models.GetDB().Where("id = ?", eventId).Preload("Users").First(updatedEvent).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": updatedEvent, "success": true})
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	authorizedUserId := c.GetUint("user")
	authorizedUser := models.GetUser(authorizedUserId)

	if authorizedUser.Role != models.Admin && authorizedUser.Role != models.SuperAdmin {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "current user does't have rights to perform action", "success": false})
		return
	}

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

	c.JSON(http.StatusOK, gin.H{"event id": *eventId, "success": true})
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
