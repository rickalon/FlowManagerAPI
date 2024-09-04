package domain

import (
	"errors"

	"github.com/rickalon/FlowManagerAPI/internal/repositories"
)

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	CreatedTime string `json:"date"`
}

func ValidateUser(user *User) error {
	var str string
	if user.Name == "" {
		str += "name should be included."
	}
	if user.Email == "" {
		str += "email should be included."
	}
	if user.Password == "" {
		str += "password should be included."
	}
	if len(str) != 0 {
		return errors.New(str)
	}
	return nil
}

func ValidateUserLogin(user *User) error {
	var str string
	if user.Name != "" {
		str += "name shouldn't be included."
	}
	if user.Email == "" {
		str += "email should be included."
	}
	if user.Password == "" {
		str += "password should be included."
	}
	if len(str) != 0 {
		return errors.New(str)
	}
	return nil
}

func CreateUser(db *repositories.PqDB, user *User) error {
	_, err := db.DB.Exec("INSERT INTO USERS(full_name,password,email) VALUES ($1,$2,$3);", user.Name, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func GetIdUser(db *repositories.PqDB, user *User) error {
	err := db.DB.QueryRow("SELECT user_id from USERS where email=$1", user.Email).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetLoginUser(db *repositories.PqDB, user *User) error {
	err := db.DB.QueryRow("SELECT user_id,full_name,password from USERS where email=$1", user.Email).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		return err
	}
	return nil
}
