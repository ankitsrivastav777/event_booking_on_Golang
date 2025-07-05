package models

import (
	"booking/rest-api/db"
	"booking/rest-api/utils"
)

type User struct {
	ID       int64
	Password string `binding:"required"`
	Email    string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Hash the password before saving it to the database
	hashedPassword, err := utils.HashPassword(u.Password)

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return err
}
