package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Period    time.Time `json:"period"`
	Address   string    `json:"address"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
}
