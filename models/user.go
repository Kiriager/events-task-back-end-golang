package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func UpdateUserRecord(updateUserData *UpdateUser, userId *uint, fromAdmin bool) error {
	userToUpdate, err := GetUser(*userId)
	if err != nil {
		return err
	}

	changePassword := false
	if updateUserData.Password != "" {
		userToUpdate.Password = updateUserData.Password
		changePassword = true
	}
	if fromAdmin {
		userToUpdate.Role = Role(updateUserData.Role)
	}

	err = userToUpdate.Validate()
	if err != nil {
		return err
	}

	if changePassword {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userToUpdate.Password), bcrypt.DefaultCost)
		userToUpdate.Password = string(hashedPassword)
	}

	err = GetDB().Updates(userToUpdate).Where("id = ?", *userId).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateMyRecord(updateUserData *UpdateMyAcc, userId *uint) error {
	me, err := GetUser(*userId)
	if err != nil {
		return err
	}
	changePassword := false

	if updateUserData.Password != "" {
		me.Password = updateUserData.Password
		changePassword = true
	}

	err = me.Validate()
	if err != nil {
		return err
	}

	if changePassword {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(me.Password), bcrypt.DefaultCost)
		me.Password = string(hashedPassword)
	}

	err = GetDB().Updates(me).Where("id = ?", *userId).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUser(userId uint) (*User, error) {

	acc := &User{}
	err := GetDB().Where("id = ?", userId).First(acc).Error

	if err != nil {
		return nil, err
	}

	//acc.Password = ""
	return acc, nil
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

func DeleteUser(userId uint) error {
	//user := GetUser(userId)

	err := GetDB().Delete(&User{}, userId).Error
	if err != nil {
		return err
	}
	return nil
}
