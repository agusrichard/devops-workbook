package repository

import (
	"fmt"
	"golang-restapi/models"
)

// CreateUser -- Create user
func CreateUser(user *models.User) {
	sqlQuery := `
		INSERT INTO users (username, password) 
		VALUES ($1, $2)
	`

	_, err := DB.Exec(sqlQuery, user.Username, user.Password)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Success to create user %v\n", *user)
}

// GetUserByUsername -- Get user by username
func GetUserByUsername(username string) models.User {
	sqlQuery := `
		SELECT * FROM users
		WHERE username = $1
	`
	rows, err := DB.Query(sqlQuery, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user models.User
	for rows.Next() {
		err = rows.Scan(
			&user.Username,
			&user.Password,
		)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return user
}
