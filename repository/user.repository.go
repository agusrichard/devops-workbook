package repository

import (
	"fmt"
	"golang-restapi/config"
	"golang-restapi/model"
)

// CreateUser -- Create user
func CreateUser(user *model.User) bool {
	sqlQuery := `
		INSERT INTO users (username, password) 
		VALUES ($1, $2);
	`

	_, err := config.DB.Exec(sqlQuery, user.Username, user.Password)
	if err != nil {
		return false
	}
	fmt.Printf("Success to create user %v\n", *user)
	return true
}

// GetUserByUsername -- Get user by username
func GetUserByUsername(username string) (model.User, error) {
	sqlQuery := `
		SELECT _id, username, password FROM users
		WHERE username = $1;
	`
	rows, err := config.DB.Query(sqlQuery, username)
	if err != nil {
		return model.User{}, err
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
			return model.User{}, err
		}
	}
	err = rows.Err()
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GetUserByID -- Get user data
func GetUserByID(userID uint64) (model.User, error) {
	sqlQuery := `
		SELECT _id, username, password FROM users
		WHERE _id = $1
	`
	rows, err := config.DB.Query(sqlQuery, userID)
	if err != nil {
		return model.User{}, err
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
			return model.User{}, err
		}
	}
	err = rows.Err()
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
