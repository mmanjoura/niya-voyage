package models

import "time"

type Car struct {
	ID           int       `json:"id"  gorm:"primary_key"`
	Tag          string    `json:"tag"`
	Title        string    `json:"title"`
	Price        string    `json:"price"`
	Location     string    `json:"location"`
	Duration     string    `json:"duration"`
	Reviews      string    `json:"reviews"`
	Seat         string    `json:"seat"`
	Type         string    `json:"type"`
	Luggage      string    `json:"luggage"`
	Animation    string    `json:"animation"`
	Transmission string    `json:"transmission"`
	Speed        string    `json:"speed"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateCar struct {
	Tag          string `json:"tag"`
	Title        string `json:"title"`
	Price        string `json:"price"`
	Location     string `json:"location"`
	Duration     string `json:"duration"`
	Reviews      string `json:"reviews"`
	Seat         string `json:"seat"`
	Type         string `json:"type"`
	Luggage      string `json:"luggage"`
	Animation    string `json:"animation"`
	Transmission string `json:"transmission"`
	Speed        string `json:"speed"`
}

type UpdateCar struct {
	Tag          string `json:"tag"`
	Title        string `json:"title"`
	Price        string `json:"price"`
	Location     string `json:"location"`
	Duration     string `json:"duration"`
	Reviews      string `json:"reviews"`
	Seat         string `json:"seat"`
	Type         string `json:"type"`
	Luggage      string `json:"luggage"`
	Animation    string `json:"animation"`
	Transmission string `json:"transmission"`
	Speed        string `json:"speed"`
}
