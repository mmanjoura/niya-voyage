package models

import "time"

type Hotel struct {
	ID              int       `json:"id"  gorm:"primary_key"`
	Category_ID     int       `json:"category_id"`
	Tag             string    `json:"tag"`
	Img             string    `json:"img"`
	Title           string    `json:"title"`
	Location        string    `json:"location"`
	Rating          float32   `json:"rating"`
	NumberOfReviews int       `json:"number_of_reviews"`
	Price           float32   `json:"price"`
	DelayAnimation  int       `json:"delay_animation"`
	City            string    `json:"city"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateHotel struct {
	Category_ID     int     `json:"category_id"`
	Tag             string  `json:"tag"`
	Img             string  `json:"img"`
	Title           string  `json:"title"`
	Location        string  `json:"location"`
	Rating          float32 `json:"rating"`
	NumberOfReviews int     `json:"number_of_reviews"`
	Price           float32 `json:"price"`
	DelayAnimation  int     `json:"delay_animation"`
	City            string  `json:"city"`
}

type UpdateHotel struct {
	Category_ID     int     `json:"category_id"`
	Tag             string  `json:"tag"`
	Img             string  `json:"img"`
	Title           string  `json:"title"`
	Location        string  `json:"location"`
	Rating          float32 `json:"rating"`
	NumberOfReviews int     `json:"number_of_reviews"`
	Price           float32 `json:"price"`
	DelayAnimation  int     `json:"delay_animation"`
	City            string  `json:"city"`
}
