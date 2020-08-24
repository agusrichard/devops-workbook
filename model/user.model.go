package model

// User Model
type User struct {
	ID        uint64 `json:"_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UUID      string `json:"uuid"`
	Confirmed bool   `json:"confirmed"`
}

// ConfirmData ---  Used in ConfirmAccount handler
type ConfirmData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
}
