package models

import (
	"time"
)

type SlideImage struct {
	ID         int       `json:"id"  gorm:"primary_key"`
	HotelID    int       `json:"HotelID"`
	ActivityID int       `json:"ActivityID"`
	CarID      int       `json:"CarID"`
	GolfID     int       `json:"GolfID"`
	RentalID   int       `json:"RentalID"`
	TourID     int       `json:"TourID"`
	Img        string    `json:"img"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateSlideImage struct {
	HotelID    int    `json:"HotelID"`
	ActivityID int    `json:"ActivityID"`
	CarID      int    `json:"CarID"`
	GolfID     int    `json:"GolfID"`
	RentalID   int    `json:"RentalID"`
	TourID     int    `json:"TourID"`
	Img        string `json:"img"`
}

type UpdateSlideImage struct {
	HotelID    int    `json:"HotelID"`
	ActivityID int    `json:"ActivityID"`
	CarID      int    `json:"CarID"`
	GolfID     int    `json:"GolfID"`
	RentalID   int    `json:"RentalID"`
	TourID     int    `json:"TourID"`
	Img        string `json:"img"`
}
