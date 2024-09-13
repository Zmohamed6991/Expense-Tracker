package models

import (
	"gorm.io/gorm"
	
)

type Expenses struct {
	gorm.Model
	ExpenseName string  `json:"expense_name"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
}
