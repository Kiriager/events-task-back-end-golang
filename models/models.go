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
	Email    string   `json:"email"`
	Password string   `json:"password,omitempty"`
	Role     Role     `json:"role"` //Role *Role can't have address
	Token    string   `json:"token,omitempty" sql:"-"`
	Events   []*Event `gorm:"many2many:user_events;"`
}

type Role string

const (
	Regular    Role = "regular"    //able to edit own profile and manage own events
	Admin      Role = "admin"      //able to create locations and events
	SuperAdmin Role = "superadmin" //admin + able to manage users profiles
)

type UserAuth struct {
	gorm.Model
	UserID   uint   `gorm:";not null;" json:"user_id"`
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
}

type RegisterUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUser struct {
	//Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
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
	Users       []*User   `gorm:"many2many:user_events;"`
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
	Events      []Event
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

type RegUserToEvent struct {
	UserId  uint `json:"userId" binding:"required"`
	Status  bool `json:"status" binding:"required"`
	EventId uint `json:"eventId" binding:"required"`
}
