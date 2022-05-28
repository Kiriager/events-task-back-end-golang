package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

/*
JWT claims struct
*/

type Token struct {
	UserId   uint
	AuthUUID string
	jwt.RegisteredClaims
}

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty" sql:"-"`
}

type UserAuth struct {
	gorm.Model
	UserID   uint   `gorm:";not null;" json:"user_id"`
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
}

type CreateUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateEvent struct {
	Title       string `json:"title" binding:"required, min=2, max=40"`
	Description string `json:"description" binding:"max=100"`
	Start       string `json:"start" binding:"required"`
	End         string `json:"end" binding:"required"`
	Location    string `json:"location" binding:"required, min=2, max=40"`
	Latitude    string `json:"latitude" binding:"required"`
	Longitude   string `json:"longitude" binding:"required"`
}

type Event struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Location    string `json:"location"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type UpdateEvent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Location    string `json:"location"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}
