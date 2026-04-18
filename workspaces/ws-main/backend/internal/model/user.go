package model

type User struct {
	ID           int64
	Username     string
	PasswordHash string
	Role         string
	Status       string
}

type Session struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}
