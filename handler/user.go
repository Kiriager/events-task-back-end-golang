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

	authorizedUserId := c.GetUint("user")
	userId, err := h.getPathParamUint(c, "userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	if *userId != authorizedUserId {
		authorizedUser := models.GetUser(authorizedUserId)
		if authorizedUser.Role != models.SuperAdmin {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("current user does't have rights to perform action"), "success": false})
			return
		}
	}
	updatedUser, err := models.UpdateUserRecord(&userUpdateData, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": updatedUser, "success": true})
}
