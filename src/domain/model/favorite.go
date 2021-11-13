package model

import "time"

type Favorite struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	PlaceID   string    `json:"place_id"`
	CreatedAd time.Time `json:"created_ad"`
	UpdatedAt time.Time `json:"updated_at"`

	User User
}
