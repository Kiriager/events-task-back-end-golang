package models

import (
	"errors"
	"os"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//Validate incoming user details...

func (account *User) Validate() error {

	if !strings.Contains(account.Email, "@") {
		return errors.New("email address is required")
	}

	if len(account.Password) < 6 {
		return errors.New("password is required")
	}

	//Email must be unique
	temp := &User{}

	//check for errors and duplicate emails
	err := GetDB().Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("connection error. please retry")
	}
	if temp.Email != "" && temp.ID != account.ID {
		return errors.New("email address already in use by another user")
	}

	return nil
}

func CreateToken(authD *UserAuth) (string, error) {

	tk := &Token{UserId: authD.UserID, AuthUUID: authD.AuthUUID}
	//tk.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 15))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	return token.SignedString([]byte(os.Getenv("token_password")))
}

func FetchAuth(token *Token) (*UserAuth, error) {

	au := &UserAuth{}
	err := GetDB().Debug().Where("user_id = ? AND auth_uuid = ?", token.UserId, token.AuthUUID).Take(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

// DeleteAuth Once a user row in the auth table
func DeleteAuth(authD *UserAuth) error {

	au := &UserAuth{}
	db := GetDB().Debug().Where("user_id = ? AND auth_uuid = ?", authD.UserID, authD.AuthUUID).Take(&au).Delete(&au)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// CreateAuth Once the user signup/login, create a row in the auth table, with a new uuid
func CreateAuth(userId uint) (*UserAuth, error) {

	au := &UserAuth{}
	uuidV4, _ := uuid.NewV4()
	au.AuthUUID = uuidV4.String() //generate a new UUID each time
	au.UserID = userId
	err := GetDB().Debug().Create(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

func Login(request *LoginRequest) (*User, error) {

	db := GetDB()
	user := &User{}
	err := db.Where("email = ?", request.Email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("email address not found")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return nil, errors.New("invalid login credentials. Please try again")
	}
	//Worked! Logged In
	user.Password = ""

	//Create JWT token

	user.Token = GenerateToken(user.ID)

	return user, nil
}

func GenerateToken(userId uint) string {
	authUser, _ := CreateAuth(userId)
	tokenString, _ := CreateToken(authUser)
	return tokenString
}

func Logout(user uint, auth string) error {

	userAuth := &UserAuth{UserID: user, AuthUUID: auth}
	err := DeleteAuth(userAuth)
	if err != nil {
		return errors.New("auth not found")
	}
	return nil
}
