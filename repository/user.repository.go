package repository

import (
	"fmt"
	"golang-restapi/model"
)

// CreateUser -- Create user
func CreateUser(user *model.User) {
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
func GetUserByUsername(username string) model.User {
	sqlQuery := `
		SELECT _id, username, password FROM users
		WHERE username = $1
	`
	rows, err := DB.Query(sqlQuery, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user model.User
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
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

// GetUserByID -- Get user data
func GetUserByID(userID uint64) model.User {
	sqlQuery := `
		SELECT _id, username, password FROM users
		WHERE _id = $1
	`
	rows, err := DB.Query(sqlQuery, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user model.User
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
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
