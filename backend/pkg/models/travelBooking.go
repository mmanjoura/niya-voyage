package models

import "time"

type TravelBooking struct {
	ID          int       `json:"id"  gorm:"primary_key"`
	Customer_ID int       `json:"customer_id"`
	Package_ID  int       `json:"package_id"`
	Travel_Date string    `json:"travel_date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateTravelBooking struct {
	Customer_ID int    `json:"customer_id"`
	Package_ID  int    `json:"package_id"`
	Travel_Date string `json:"travel_date"  binding:"required"`
	Status      string `json:"status"  binding:"required"`
}

type UpdateTravelBooking struct {
	Customer_ID int    `json:"customer_id"`
	Package_ID  int    `json:"package_id"`
	Travel_Date string `json:"travel_date"`
	Status      string `json:"status"`
}
