package models

import (
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

func UserExists(email string) (bool, error) {
	rows, err := database.DB.Query("SELECT * FROM users WHERE email = ?", email)

	if err != nil {
		return false, err
	}

	count := 0
	for rows.Next() {
		count += 1
	}

	defer rows.Close()

	if err := rows.Err(); err != nil {
		return false, err
	}

	return count == 1, err
}
