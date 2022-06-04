package models

import (
	"time"

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

type RegisterEvent struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Start       time.Time `json:"start" binding:"required"`
	End         time.Time `json:"end" binding:"required"`
	LocationID  uint      `json:"locationid" binding:"required"`
}

type Event struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	LocationID  uint      `json:"locationid"`
}

type UpdateEvent struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	LocationID  uint      `json:"location"`
}

type Location struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type RegisterLocation struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude" binding:"required"`
	Longitude   float64 `json:"longitude" binding:"required"`
}
