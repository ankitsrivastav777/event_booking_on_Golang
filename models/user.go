package models

import (
	"booking/rest-api/db"
	"booking/rest-api/utils"
	"errors"
	"fmt"
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
	if err != nil {
		return err
	}
	fmt.Println(hashedPassword, "User saved successfully")
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

func (u *User) ValidatePassword() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var hashedPassword string
	err := row.Scan(&u.ID, &hashedPassword)
	// Retrieve the hashed password from the database
	if err != nil {
		return errors.New("User not found")
	}

	// Compare the provided password with the hashed password from the database
	err = utils.CheckPasswordHash(u.Password, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}
