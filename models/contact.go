package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Company  string `json:"company"`
	TelText  string `json:"telText"`
	TelValue string `json:"telValue"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
