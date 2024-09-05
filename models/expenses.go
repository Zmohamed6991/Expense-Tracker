package models

import (
	
	"gorm.io/gorm"
)

type Expenses struct {
	gorm.Model
	ID int `json:"id" gorm:"primary_key"`
	ExpenseName string `json:"expense_name"`
	Amount float64 `json:"amount"`
	CategoryID int `json:"category_id"`
	Categories Categories `gorm:"foreign_key:CategoryID"`
	
	
	



}