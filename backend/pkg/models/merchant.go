package models

import "time"

type Merchant struct {
	ID                  int       `json:"id"  gorm:"primary_key"`
	BusinessName        string    `json:"business_name "`
	UserName            string    `json:"user_name"`
	FirstName           string    `json:"firs_tname"`
	LastName            string    `json:"last_name"`
	Email               string    `json:"email"`
	PhoneNumber         string    `json:"phone_number"`
	BirthDate           string    `json:"birth_date"`
	About               string    `json:"about"`
	LocationInformation Location  `json:"location_information" gorm:"foreignKey:MerchantID"`
	ChangePassword      Password  `json:"change_password" gorm:"foreignKey:MerchantID"`
	CreatedAt           time.Time `json:"created_at"  gorm:"autoCreateTime"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateMerchant struct {
	BusinessName string `json:"business_name "`
	UserName     string `json:"user_name"`
	FirstName    string `json:"firs_tname"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	BirthDate    string `json:"birth_date"`
	About        string `json:"about"`
}

type UpdateMerchant struct {
	BusinessName string `json:"business_name "`
	UserName     string `json:"user_name"`
	FirstName    string `json:"firs_tname"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	BirthDate    string `json:"birth_date"`
	About        string `json:"about"`
}

type Location struct {
	ID           int    `json:"id"  gorm:"primary_key"`
	MerchantID   int    `json:"MerchantID"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}

type Password struct {
	ID               int    `json:"id"  gorm:"primary_key"`
	MerchantID       int    `json:"MerchantID"`
	CurrentPassword  string `json:"current_password"`
	NewPassword      string `json:"new_password"`
	NewPasswordAgain string `json:"new_password_again"`
}

type CreateLocation struct {
	MerchantID   int    `json:"MerchantID"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}

type CreatePassword struct {
	MerchantID       int    `json:"MerchantID"`
	CurrentPassword  string `json:"current_password"`
	NewPassword      string `json:"new_password"`
	NewPasswordAgain string `json:"new_password_again"`
}
type UpdateLocation struct {
	MerchantID   int    `json:"MerchantID"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}

type UpdatePassword struct {
	MerchantID       int    `json:"MerchantID"`
	CurrentPassword  string `json:"current_password"`
	NewPassword      string `json:"new_password"`
	NewPasswordAgain string `json:"new_password_again"`
}
