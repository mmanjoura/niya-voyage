package models

import "time"

type TourPackage struct {
	ID           int       `json:"id"  gorm:"primary_key"`
	Package_Name string    `json:"package_Name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Itinerary    string    `json:"itinerary"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateTourPackage struct {
	Package_Name string  `json:"package_Name"  binding:"required"`
	Description  string  `json:"description"  binding:"required"`
	Price        float64 `json:"price"  binding:"required"`
	Itinerary    string  `json:"itinerary"  binding:"required"`
}

type UpdateTourPackage struct {
	Package_Name string  `json:"package_Name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Itinerary    string  `json:"itinerary"`
}
