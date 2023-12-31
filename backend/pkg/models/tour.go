package models

import "time"

type Tour struct {
	ID            int            `json:"id"  gorm:"primary_key"`
	Tag           string         `json:"tag"`
	Title         string         `json:"title"`
	Price         string         `json:"price"`
	Location      string         `json:"location"`
	Duration      string         `json:"duration"`
	Reviews       string         `json:"reviews"`
	TourType      string         `json:"tourType"`
	Animation     string         `json:"animation"`
	SlideImages   []SlideImage   `json:"SlideImages" gorm:"foreignKey:TourID"`
	GalleryImages []GalleryImage `json:"GalleryImages" gorm:"foreignKey:HotelID"`
	SlideImg      []string       `json:"slide_img"`
	GalleryImg    []string       `json:"gallery_img"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateTour struct {
	Tag           string         `json:"tag"`
	Title         string         `json:"title"`
	Price         string         `json:"price"`
	Location      string         `json:"location"`
	Duration      string         `json:"duration"`
	Reviews       string         `json:"reviews"`
	TourType      string         `json:"tourType"`
	Animation     string         `json:"animation"`
	SlideImages   []SlideImage   `json:"SlideImages" gorm:"foreignKey:TourID"`
	GalleryImages []GalleryImage `json:"GalleryImages" gorm:"foreignKey:HotelID"`
	SlideImg      []string       `json:"slide_img"`
	GalleryImg    []string       `json:"gallery_img"`
}

type UpdateTour struct {
	Tag           string         `json:"tag"`
	Title         string         `json:"title"`
	Price         string         `json:"price"`
	Location      string         `json:"location"`
	Duration      string         `json:"duration"`
	Reviews       string         `json:"reviews"`
	TourType      string         `json:"tourType"`
	Animation     string         `json:"animation"`
	SlideImages   []SlideImage   `json:"SlideImages" gorm:"foreignKey:TourID"`
	GalleryImages []GalleryImage `json:"GalleryImages" gorm:"foreignKey:HotelID"`
	SlideImg      []string       `json:"slide_img"`
	GalleryImg    []string       `json:"gallery_img"`
}
