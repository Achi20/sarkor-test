package entity

import "time"

type User struct {
	UserID    int       `json:"user_id"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Name      string    `json:"is_fax"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}
