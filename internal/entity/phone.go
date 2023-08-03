package entity

import "time"

type Phone struct {
	PhoneID     int       `json:"phone_id"`
	Phone       string    `json:"phone"`
	Description string    `json:"description"`
	IsFax       bool      `json:"is_fax"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}
