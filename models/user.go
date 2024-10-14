package models

import (
	"errors"
	"example/mysql-api/database"
	"example/mysql-api/utils"
)

type User struct {
	ID       int64
	Email    string `bindings:"required"`
	Password string `bindings:"required"`
}

func (user *User) Save() error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := database.DB.Exec("INSERT INTO users(email, password) VALUES (?, ?)", user.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	user.ID = id
	return err
}

func UserExists(email string) bool {
	var userEmail string
	row := database.DB.QueryRow("SELECT email FROM users WHERE email = ?", email)

	err := row.Scan(&userEmail)

	return err == nil
}

func (user *User) ValidateCredentials() error {
	row := database.DB.QueryRow("SELECT password FROM users WHERE email = ?", user.Email)

	var hashedPassword string
	err := row.Scan(&hashedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(hashedPassword, user.Password)

	if !passwordIsValid {
		return errors.New("Invalid Credentials")
	}

	return nil
}
