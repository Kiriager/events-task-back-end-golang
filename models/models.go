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
	Role     *Role  `json:"role"`
	Token    string `json:"token,omitempty" sql:"-"`
}

type Role string

const (
	Regular Role = "regular"
	Admin   Role = "admin"
)

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
	LocationId  uint      `json:"locationid" binding:"required"`
}

type Event struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	LocationId  uint      `json:"locationid"`
	Location    Location  `gorm:"ForeignKey:LocationId"`
}

type UpdateEvent struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	LocationId  uint      `json:"locationid"`
}

type Location struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Events      []Event //`gorm:"ForeignKey:LocationId"`
}

type RegisterLocation struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude" binding:"required"`
	Longitude   float64 `json:"longitude" binding:"required"`
}

type UpdateLocation struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
