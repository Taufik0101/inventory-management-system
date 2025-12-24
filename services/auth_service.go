package services

import (
	"errors"
	"inventory-management-system/models"
	"inventory-management-system/utils"

	"golang.org/x/crypto/bcrypt"
)

func Register(username, password string) error {
	for _, u := range models.Users {
		if u.Username == username {
			return errors.New("user already exists")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := models.User{
		ID:       models.NextUserID,
		Username: username,
		Password: string(hashedPassword),
	}
	models.Users = append(models.Users, newUser)
	models.NextUserID++
	return nil
}

func Login(username, password string) (string, error) {
	var user models.User
	found := false
	for _, u := range models.Users {
		if u.Username == username {
			user = u
			found = true
			break
		}
	}

	if !found {
		return "", errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(username)
}
