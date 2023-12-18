package models

import (
	"time"
)

type SlideImage struct {
	ID          int       `json:"id"  gorm:"primary_key"`
	HotelID     int       `json:"HotelID"`
	Activity_ID int       `json:"activity_ID"`
	Car_ID      int       `json:"car_ID"`
	Golf_ID     int       `json:"golf_ID"`
	Rental_ID   int       `json:"rental_ID"`
	Tour_ID     int       `json:"tour_ID"`
	SlideImg    string    `json:"slideImg"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateSlideImage struct {
	HotelID     int    `json:"HotelID"`
	Activity_ID int    `json:"activity_ID"`
	Car_ID      int    `json:"car_ID"`
	Golf_ID     int    `json:"golf_ID"`
	Rental_ID   int    `json:"rental_ID"`
	Tour_ID     int    `json:"tour_ID"`
	SlideImg    string `json:"slideImg"`
}

type UpdateSlideImage struct {
	HotelID     int    `json:"HotelID"`
	Activity_ID int    `json:"activity_ID"`
	Car_ID      int    `json:"car_ID"`
	Golf_ID     int    `json:"golf_ID"`
	Rental_ID   int    `json:"rental_ID"`
	Tour_ID     int    `json:"tour_ID"`
	SlideImg    string `json:"slideImg"`
}
