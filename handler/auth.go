package handler

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	registerUser := models.RegisterUser{}
	if err := c.ShouldBindJSON(&registerUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	user, err := models.CreateUser(&registerUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	user.Password = ""
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
	user.Password = ""
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
	user, err := models.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"user": user, "success": true})
}
