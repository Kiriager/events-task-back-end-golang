package handler

import (
	"errors"
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateUser(c *gin.Context) { //not finished
	userUpdateData := models.UpdateUser{}

	err := c.ShouldBindJSON(&userUpdateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	updateUserId, err := h.getPathParamUint(c, "userId") //user should be updated
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	authorizedUserId := c.GetUint("user")
	authorizedUser := models.GetUser(authorizedUserId)
	updatedUser := models.GetUser(*updateUserId)

	if authorizedUser.Role != models.SuperAdmin {
		//manage account from superadmin
		err = models.UpdateUserRecord(&userUpdateData, updateUserId, true)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
			return
		}
	} else if *updateUserId == authorizedUserId {
		err = models.UpdateUserRecord(&userUpdateData, updateUserId, false)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": errors.New("current user does't have rights to perform action"), "success": false})
		return
	}

	//updatedUser, err := models.UpdateUserRecord(&userUpdateData, updateUserId)

	c.JSON(http.StatusOK, gin.H{"user": updatedUser, "success": true})
}

func (h *Handler) manageUserEvent(c *gin.Context) {
	request := models.RegUserToEvent{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	authorizedUserId := c.GetUint("user")

	regUserId := request.UserId
	eventId := request.EventId
	status := request.Status

	event, err := models.GetEvent(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	authorizedUser := models.GetUser(authorizedUserId)
	regUser := models.GetUser(regUserId)

	if authorizedUserId == regUserId { //reg from own acc
		if status {
			err = models.GetDB().Model(&authorizedUser).Association("Events").Append(event)
		} else {
			err = models.GetDB().Model(&authorizedUser).Association("Events").Delete(event)
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
			return
		}
		err = models.GetDB().Where("id = ?", authorizedUserId).Preload("Events").First(authorizedUser).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user's events": authorizedUser.Events, "success": true})
	}

	if authorizedUser.Role != models.SuperAdmin {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "current user does't have rights to perform action", "success": false})
		return
	}
	//reg user to event from super admin

	if status {
		err = models.GetDB().Model(&regUser).Association("Events").Append(event)
	} else {
		err = models.GetDB().Model(&regUser).Association("Events").Delete(event)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	err = models.GetDB().Where("id = ?", regUserId).Preload("Events").First(regUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user's events": regUser.Events, "success": true})
}

func (h *Handler) ShowUserEvents(c *gin.Context) {
	authorizedUserId := c.GetUint("user")
	authorizedUser := models.GetUser(authorizedUserId)

	err := models.GetDB().Where("id = ?", authorizedUserId).Preload("Events").First(authorizedUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user's events": authorizedUser.Events, "success": true})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	authorizedUserId := c.GetUint("user")
	authorizedUser := models.GetUser(authorizedUserId)

	if authorizedUser.Role != models.SuperAdmin {
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
