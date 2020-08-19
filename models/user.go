package models

// User Model
type User struct {
	ID       uint64 `json:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
