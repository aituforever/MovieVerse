package models

import "gorm.io/datatypes"

type Movie struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Director    string         `json:"director"`
	Country     string         `json:"country"`
	Genres      datatypes.JSON `json:"genres" gorm:"type:jsonb"` // Change this to jsonb
	ReleaseYear int            `json:"release_year"`
	Description string         `json:"description"`
}
