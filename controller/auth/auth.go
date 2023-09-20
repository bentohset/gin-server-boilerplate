package authController

import (
	"example/gin-webserver/database"
	"example/gin-webserver/model"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

// adds user to the database
func CreateUser(user model.User) (model.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// []byte is byte slice
	if err != nil {
		return model.User{}, err
	}

	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	err = database.Database.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// validate password for given user
func ValidatePassword(user model.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (model.User, error) {
	var user model.User
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (model.User, error) {
	var user model.User
	// entries associated to user prepopulated
	err := database.Database.Preload("Entries").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
