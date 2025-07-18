package models

import (
	"example.com/investment-calulator/db"
	"example.com/investment-calulator/utils"
)

type User struct {
	ID 			int64
	Email 		string `binding:"required"`
	Password 	string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil { return err }

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil { return err }

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil { return err }

	userID, err := result.LastInsertId()
	u.ID = userID

	return err
}