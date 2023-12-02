package models

import "time"

type Destination struct {
	ID             int       `json:"id"  gorm:"primary_key"`
	ColClass       string    `json:"col_Class"`
	Title          string    `json:"title"`
	Location       string    `json:"location"`
	Travellers     string    `json:"travellers"`
	HoverText      string    `json:"hover_text"`
	Img            string    `json:"img"`
	City           string    `json:"city"`
	Properties     string    `json:"properties"`
	Region         string    `json:"region"`
	DelayAnimation int       `json:"delay_animation"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateDestination struct {
	ColClass       string `json:"col_Class"`
	Title          string `json:"title"`
	Location       string `json:"location"`
	Travellers     string `json:"travellers"`
	HoverText      string `json:"hover_text"`
	Img            string `json:"img"`
	City           string `json:"city"`
	Properties     string `json:"properties"`
	Region         string `json:"region"`
	DelayAnimation int    `json:"delay_animation"`
}

type UpdateDestination struct {
	ColClass       string `json:"col_Class"`
	Title          string `json:"title"`
	Location       string `json:"location"`
	Travellers     string `json:"travellers"`
	HoverText      string `json:"hover_text"`
	Img            string `json:"img"`
	City           string `json:"city"`
	Properties     string `json:"properties"`
	Region         string `json:"region"`
	DelayAnimation int    `json:"delay_animation"`
}
