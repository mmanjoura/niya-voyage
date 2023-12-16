package models

import "time"

type Rental struct {
	ID        int       `json:"id"  gorm:"primary_key"`
	Tag       string    `json:"tag"`
	Title     string    `json:"title"`
	Price     string    `json:"price"`
	Location  string    `json:"location"`
	Duration  string    `json:"duration"`
	Reviews   string    `json:"reviews"`
	Ratings   string    `json:"ratings"`
	Guest     string    `json:"guest"`
	Bedroom   string    `json:"bedroom"`
	Bed       string    `json:"bed"`
	Animation string    `json:"animation"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateRental struct {
	Tag       string `json:"tag"`
	Title     string `json:"title"`
	Price     string `json:"price"`
	Location  string `json:"location"`
	Duration  string `json:"duration"`
	Reviews   string `json:"reviews"`
	Ratings   string `json:"ratings"`
	Guest     string `json:"guest"`
	Bedroom   string `json:"bedroom"`
	Bed       string `json:"bed"`
	Animation string `json:"animation"`
}

type UpdateRental struct {
	Tag       string `json:"tag"`
	Title     string `json:"title"`
	Price     string `json:"price"`
	Location  string `json:"location"`
	Duration  string `json:"duration"`
	Reviews   string `json:"reviews"`
	Ratings   string `json:"ratings"`
	Guest     string `json:"guest"`
	Bedroom   string `json:"bedroom"`
	Bed       string `json:"bed"`
	Animation string `json:"animation"`
}
