package models

import (
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	ID int `json:"id" gorm:"primary_key"`
	CategoryName string `json:"category_name"`
	Categories []Categories `gorm:"foreign_key:CategoryID"`


}