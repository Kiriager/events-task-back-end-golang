package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func UpdateUserRecord(updateUserData *UpdateUser, userId *uint, fromAdmin bool) error {
	userToUpdate := GetUser(*userId)
	if updateUserData.Password != "" {
		userToUpdate.Password = updateUserData.Password
	}
	if fromAdmin {
		userToUpdate.Role = updateUserData.Role
	}

	err := userToUpdate.Validate()
	if err != nil {
		return err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userToUpdate.Password), bcrypt.DefaultCost)
	userToUpdate.Password = string(hashedPassword)

	err = GetDB().Updates(userToUpdate).Where("id = ?", *userId).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUser(userId uint) *User {

	acc := &User{}
	GetDB().Where("id = ?", userId).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}

func CreateUser(request *RegisterUser) (*User, error) {
	user := &User{Email: request.Email, Password: request.Password, Role: Regular}

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return nil, errors.New("failed to create user, connection error")
	}

	//Create new JWT token for the newly registered user

	user.Token = GenerateToken(user.ID)

	user.Password = "" //delete password

	return user, nil
}
