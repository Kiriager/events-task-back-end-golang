package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	userCreate := models.CreateUser{}
	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	user, err := models.Create(&userCreate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "success": true})
}

func (h *Handler) SignIn(c *gin.Context) {
	loginRequest := models.LoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	user, err := models.Login(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "success": true})
}

func (h *Handler) Logout(c *gin.Context) {
	userId := c.GetUint("user")
	authUUID := c.GetString("auth")
	err := models.Logout(userId, authUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out", "success": true})
}

func (h *Handler) MyAcc(c *gin.Context) {
	userId := c.GetUint("user")
	user := models.GetUser(userId)
	c.JSON(http.StatusOK, gin.H{"user": user, "success": true})
}
