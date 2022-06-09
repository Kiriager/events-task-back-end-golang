package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func UpdateUserRecord(updateUserData *UpdateUser, userId *uint) (*User, error) {
	//authorizedUserID = 2

	userToUpdate := GetUser(*userId)

	userToUpdate.Role = updateUserData.Role
	err := userToUpdate.Validate()
	if err != nil {
		return nil, err
	}

	err = GetDB().Updates(userToUpdate).Where("id = ?", *userId).Error
	if err != nil {
		return nil, err
	}

	/*
		userToUpdate.UpdateUserFields(updateuserData)

		err = userToUpdate.Validate()
		if err != nil {
			return nil, err
		}

		err = GetDB().Updates(userToUpdate).Where("id = ?", *userId).Error
		if err != nil {
			return nil, err
		}

		updatedUserRecord, err := GetUser(*userId)
		if err != nil {
			return nil, err
		}
	*/
	//return updatedUserRecord, nil
	return userToUpdate, nil
}

func GetUser(u uint) *User {

	acc := &User{}
	GetDB().Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}

func CreateUser(request *RegisterUser) (*User, error) {
	role := Regular
	user := &User{Email: request.Email, Password: request.Password, Role: Role("regular")}
	user.Role = role //
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
