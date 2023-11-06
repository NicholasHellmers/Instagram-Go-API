package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Message string `json:"message" gorm:"text;not null; default:null"`
}
