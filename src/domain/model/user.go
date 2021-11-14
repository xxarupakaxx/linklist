package model

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	LineUserID string    `json:"line_user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Favorite []Favorite  `json:"favorite"`
}
