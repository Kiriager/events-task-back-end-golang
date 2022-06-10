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
	authorizedUser, err := models.GetUser(authorizedUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	if authorizedUser.Role != models.SuperAdmin {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "current user does't have rights to perform action", "success": false})
		return
	}

	//manage account from superadmin
	err = models.UpdateUserRecord(&userUpdateData, updateUserId, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	updatedUser, err := models.GetUser(*updateUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	updatedUser.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": updatedUser, "success": true})
}

func (h *Handler) manageUserEvent(c *gin.Context) { //reg or dereg user from event
	request := models.RegUserToEvent{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	authorizedUserId := c.GetUint("user")
	authorizedUser, err := models.GetUser(authorizedUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	if authorizedUser.Role != models.SuperAdmin {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "current user does't have rights to perform action", "success": false})
		return
	}

	regUserId, err := h.getPathParamUint(c, "userId") //user should be updated
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	eventId := request.EventId
	status := request.Status

	event, err := models.GetEvent(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	regUser, err := models.GetUser(*regUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	if status == "reg" {
		err = models.GetDB().Model(&regUser).Association("Events").Append(event)
	} else if status == "dereg" {
		err = models.GetDB().Model(&regUser).Association("Events").Delete(event)
	} else {
		err = errors.New("unknown value of status")
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
	authorizedUser, err := models.GetUser(authorizedUserId)
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

func (h *Handler) DeleteUser(c *gin.Context) {
	authorizedUserId := c.GetUint("user")
	authorizedUser, err := models.GetUser(authorizedUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	if authorizedUser.Role != models.SuperAdmin {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "current user does't have rights to perform action", "success": false})
		return
	}

	userId, err := h.getPathParamUint(c, "userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	err = models.DeleteUser(*userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userId": *userId, "success": true})
}

func (h *Handler) ShowUser(c *gin.Context) {
	userId, err := h.getPathParamUint(c, "userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	user, err := models.GetUser(*userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"user": user, "success": true})
}

func (h *Handler) ShowMyEvents(c *gin.Context) {
	authorizedUserId := c.GetUint("user")
	authorizedUser, err := models.GetUser(authorizedUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	err = models.GetDB().Where("id = ?", authorizedUserId).Preload("Events").First(authorizedUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userEvents": authorizedUser.Events, "success": true})
}

func (h *Handler) ManageMyEvent(c *gin.Context) {
	request := models.RegUserToEvent{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	authorizedUserId := c.GetUint("user")

	eventId := request.EventId
	status := request.Status

	event, err := models.GetEvent(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	authorizedUser, err := models.GetUser(authorizedUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	if status == "reg" {
		err = models.GetDB().Model(&authorizedUser).Association("Events").Append(event)
	} else if status == "dereg" {
		err = models.GetDB().Model(&authorizedUser).Association("Events").Delete(event)
	} else {
		err = errors.New("unknown value of status")
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

	c.JSON(http.StatusOK, gin.H{"userEvents": authorizedUser.Events, "success": true})
}

func (h *Handler) UpdateMyProfile(c *gin.Context) { //not finished
	userUpdateData := models.UpdateMyAcc{}

	err := c.ShouldBindJSON(&userUpdateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	authorizedUserId := c.GetUint("user")

	err = models.UpdateMyRecord(&userUpdateData, &authorizedUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	authorizedUser, err := models.GetUser(authorizedUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	authorizedUser.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": authorizedUser, "success": true})
}
