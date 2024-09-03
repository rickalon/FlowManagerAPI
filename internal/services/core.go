package services

import (
	"errors"

	"github.com/rickalon/FlowManagerAPI/internal/domain"
	"github.com/rickalon/FlowManagerAPI/internal/repositories"
)

func ValidateUser(user *domain.User) error {
	if user.Name == "" {
		return errors.New("name should be included")
	}
	if user.Email == "" {
		return errors.New("email should be included")
	}
	if user.Password == "" {
		return errors.New("password should be included")
	}
	return nil
}

func CreateUser(db *repositories.PqDB, user *domain.User) error {
	_, err := db.DB.Exec("INSERT INTO USERS(full_name,password,email) VALUES ($1,$2,$3);", user.Name, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}
