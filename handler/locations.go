package handler

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddLocation(c *gin.Context) {
	authorizedUserId := c.GetUint("user")
	authorizedUser := models.GetUser(authorizedUserId)

	if authorizedUser.Role != models.Admin && authorizedUser.Role != models.SuperAdmin {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "current user does't have rights to perform action", "success": false})
		return
	}

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

func (h *Handler) ShowLocation(c *gin.Context) {
	locationId, err := h.getPathParamUint(c, "locationId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	//add if id == 0 or there is no id respond all events
	location, err := models.GetLocation(*locationId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"location": location, "success": true})
}

func (h *Handler) ShowAllLocations(c *gin.Context) {
	allLocations, err := models.FindAllLocations()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event": allLocations, "success": true})
}

func (h *Handler) UpdateLocation(c *gin.Context) {
	authorizedUserId := c.GetUint("user")
	authorizedUser := models.GetUser(authorizedUserId)

	if authorizedUser.Role != models.Admin && authorizedUser.Role != models.SuperAdmin {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "current user does't have rights to perform action", "success": false})
		return
	}

	locationUpdateData := models.UpdateLocation{}

	err := c.ShouldBindJSON(&locationUpdateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	locationId, err := h.getPathParamUint(c, "locationId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	updatedlocation, err := models.UpdateLocationRecord(&locationUpdateData, locationId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"location": updatedlocation, "success": true})
}

func (h *Handler) DeleteLocation(c *gin.Context) {
	authorizedUserId := c.GetUint("user")
	authorizedUser := models.GetUser(authorizedUserId)

	if authorizedUser.Role != models.Admin && authorizedUser.Role != models.SuperAdmin {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "current user does't have rights to perform action", "success": false})
		return
	}

	locationId, err := h.getPathParamUint(c, "locationId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	err = models.DeleteLocation(*locationId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
