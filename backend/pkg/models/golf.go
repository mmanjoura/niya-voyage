package models

import "time"

type Golf struct {
	ID            int            `json:"id"  gorm:"primary_key"`
	Tag           string         `json:"tag"`
	Title         string         `json:"title"`
	Name          string         `json:"name"`
	Price         string         `json:"price"`
	Location      string         `json:"location"`
	Duration      string         `json:"duration"`
	Reviews       string         `json:"reviews"`
	Ratings       string         `json:"ratings"`
	Holes         string         `json:"holes"`
	Animation     string         `json:"animation"`
	SlideImages   []SlideImage   `json:"SlideImages" gorm:"foreignKey:GolfID"`
	GalleryImages []GalleryImage `json:"GalleryImages" gorm:"foreignKey:HotelID"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateGolf struct {
	Tag           string         `json:"tag"`
	Title         string         `json:"title"`
	Name          string         `json:"name"`
	Price         string         `json:"price"`
	Location      string         `json:"location"`
	Duration      string         `json:"duration"`
	Reviews       string         `json:"reviews"`
	Ratings       string         `json:"ratings"`
	Holes         string         `json:"holes"`
	Animation     string         `json:"animation"`
	SlideImages   []SlideImage   `json:"SlideImages" gorm:"foreignKey:GolfID"`
	GalleryImages []GalleryImage `json:"GalleryImages" gorm:"foreignKey:HotelID"`
}

type UpdateGolf struct {
	Tag           string         `json:"tag"`
	Title         string         `json:"title"`
	Name          string         `json:"name"`
	Price         string         `json:"price"`
	Location      string         `json:"location"`
	Duration      string         `json:"duration"`
	Reviews       string         `json:"reviews"`
	Ratings       string         `json:"ratings"`
	Holes         string         `json:"holes"`
	Animation     string         `json:"animation"`
	SlideImages   []SlideImage   `json:"SlideImages" gorm:"foreignKey:GolfID"`
	GalleryImages []GalleryImage `json:"GalleryImages" gorm:"foreignKey:HotelID"`
}

func (Golf) TableName() string {
	return "Golfs"
}
