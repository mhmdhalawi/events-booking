package models

import (
	"github.com/mhmdhalawi/events-booking/db"
	"github.com/mhmdhalawi/events-booking/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// save user to database
func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {

		return err
	}

	defer stmt.Close()

	hashedPassword := utils.HashPassword(u.Password)

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = id

	return err

}

// find user by email
func (u *User) FindByEmail() (*User, error) {
	query := `SELECT id, email, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var user User

	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}
	return &user, err
}
