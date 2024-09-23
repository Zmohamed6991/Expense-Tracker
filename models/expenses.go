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

type Salary struct {
	gorm.Model
		MonthlySalary   float64 `json:"monthly_salary"`
		RemainingSalary float64 `json:"remaining_salary"`
	}

type Total struct {
    ID             uint      `gorm:"primaryKey"`
    ExpenseAmounts []float64 `gorm:"type:json"`  // Ensure that the database stores this as JSON or equivalent
    TotalExpenses  float64   // Field to store the sum of expenses
}

