package models

import "time"

type Activity struct {
	ID        int       `json:"id"  gorm:"primary_key"`
	Tag       string    `json:"tag"`
	Title     string    `json:"title"`
	Price     string    `json:"price"`
	Location  string    `json:"location"`
	Duration  string    `json:"duration"`
	Reviews   string    `json:"reviews"`
	Ratings   string    `json:"ratings"`
	Animation string    `json:"animation"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateActivity struct {
	Tag       string `json:"tag"`
	Title     string `json:"title"`
	Price     string `json:"price"`
	Location  string `json:"location"`
	Duration  string `json:"duration"`
	Reviews   string `json:"reviews"`
	Ratings   string `json:"ratings"`
	Animation string `json:"animation"`
}

type UpdateActivity struct {
	Tag       string `json:"tag"`
	Title     string `json:"title"`
	Price     string `json:"price"`
	Location  string `json:"location"`
	Duration  string `json:"duration"`
	Reviews   string `json:"reviews"`
	Ratings   string `json:"ratings"`
	Animation string `json:"animation"`
}