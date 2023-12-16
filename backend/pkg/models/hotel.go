package models

import "time"

type Hotel struct {
	ID          int          `json:"id"  gorm:"primary_key"`
	Category_ID int          `json:"category_id"`
	Tag         string       `json:"tag"`
	Img         string       `json:"img"`
	Title       string       `json:"title"`
	Location    string       `json:"location"`
	Ratings     float32      `json:"ratings"`
	Reviews     int          `json:"reviews"`
	Price       float32      `json:"price"`
	Animation   int          `json:"animation"`
	City        string       `json:"city"`
	HotelImages []HotelImage `json:"hotel_Images" gorm:"foreignKey:Hotel_ID"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateHotel struct {
	Category_ID int     `json:"category_id"`
	Tag         string  `json:"tag"`
	Img         string  `json:"img"`
	Title       string  `json:"title"`
	Location    string  `json:"location"`
	Ratings     float32 `json:"ratings"`
	Reviews     int     `json:"reviews"`
	Price       float32 `json:"price"`
	Animation   int     `json:"animation"`
	City        string  `json:"city"`
}

type UpdateHotel struct {
	Category_ID int     `json:"category_id"`
	Tag         string  `json:"tag"`
	Img         string  `json:"img"`
	Title       string  `json:"title"`
	Location    string  `json:"location"`
	Ratings     float32 `json:"ratings"`
	Reviews     int     `json:"reviews"`
	Price       float32 `json:"price"`
	Animation   int     `json:"animation"`
	City        string  `json:"city"`
}