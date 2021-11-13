package model

import "time"

type Place struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	PlaceID   string    `json:"place_id"`
	Address   string    `json:"address"`
	URL       string    `json:"url"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Favorite []Favorite
}
