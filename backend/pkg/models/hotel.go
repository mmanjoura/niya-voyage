package models

import (
	"time"
)

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
	SlideImages []SlideImage `json:"SlideImages" gorm:"foreignKey:HotelID"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateHotel struct {
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
	SlideImages []SlideImage `json:"SlideImages" gorm:"foreignKey:HotelID; references:ID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type UpdateHotel struct {
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
	SlideImages []SlideImage `json:"slide_images" gorm:"foreignKey:HotelID; references:ID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
