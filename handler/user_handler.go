package handler

import (
	"example/connecting/config"
	"example/connecting/models"

	"net/http"

	"github.com/gin-gonic/gin"
)


func UserSalary(c *gin.Context) {
	var Salary models.Salary

	err := c.BindJSON(&Salary)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": "invalid data"})
		return
	}

	Salary.RemainingSalary = Salary.MonthlySalary

	if err := config.DB.Create(&Salary); err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusCreated, Salary)
}

func GetAllExpense(c *gin.Context) {

	var expenses []models.Expenses

	err := config.DB.Find(&expenses)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	if len(expenses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No expenses available"})
		return
	
	}

	c.JSON(http.StatusCreated, expenses)

	var Salary models.Salary
	if err := config.DB.First(&Salary).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "unable to retrieve monthly salary"})
	}

	c.JSON(http.StatusFound, Salary)	
}

func GetExpenseByID(c *gin.Context) {
	id := c.Param("id")

	err := c.BindJSON(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	var expense models.Expenses
	if err := config.DB.First(&expense, id); err != nil {
		c.JSON(http.StatusNotFound, err)
	}
}

func CreateExpense(c *gin.Context) {
	var Salary models.Salary
	if err := config.DB.First(&Salary).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Salary not found"})
		return
	}

	var AddExpense models.Expenses
	if err := c.BindJSON(&AddExpense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	
	if AddExpense.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if AddExpense.ExpenseName == "" || AddExpense.Category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error":"no input"})
		return
	}

	if Salary.MonthlySalary == 0 || Salary.RemainingSalary == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No remaining salary available, input salary"})
		return
	}

	var Total models.Total
	if err := config.DB.Create(&Total).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error creating or fetching total"})
		return
	}

	// Add the new expense to the list of expenses
	Total.ExpenseAmounts = append(Total.ExpenseAmounts, AddExpense.Amount)

	var total float64
	for _, expense := range Total.ExpenseAmounts {
		total += expense
	}

	// Check if the expense exceeds the remaining salary
	if total > Salary.RemainingSalary {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Expense exceeds remaining salary"})
		return
	}

	// Add the new expense to the Expenses table
	if err := config.DB.Create(&AddExpense).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error adding expense"})
		return
	}

	// Update the remaining salary by subtracting the expense amount
	Salary.RemainingSalary -= AddExpense.Amount
	if err := config.DB.Save(&Salary).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error updating remaining salary"})
		return
	}

	// Save the updated total
	if err := config.DB.Save(&Total).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error saving total"})
		return
	}

	// Return the added expense, total expenses, and remaining salary
	c.JSON(http.StatusCreated, gin.H{
		"Expense added":            AddExpense,
		"Total amount of expenses": Total.ExpenseAmounts,
		"Remaining salary":         Salary.RemainingSalary,
	})
}

func UpdateAmount(c *gin.Context) {
    // Get the expense ID
    id := c.Param("id")

    // Create a struct to bind the input
    var input struct {
        Amount float64 `json:"amount"`
    }

    // Bind the input JSON
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
        return
    }

	var expenses []models.Expenses

	err := config.DB.Find(&expenses)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	if len(expenses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No expenses available to update"})
		return
	
	}

    // Find the expense we are updating
    var expense models.Expenses
    if err := config.DB.First(&expense, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
        return
    }

    // Find the salary
    var salary models.Salary
    if err := config.DB.First(&salary).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Salary not found"})
        return
    }

    // Calculate the difference between the old and new expense amount
    expenseDifference := input.Amount - expense.Amount

    // Check if the new expense exceeds the remaining salary
    if expenseDifference > salary.RemainingSalary {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Expense exceeds remaining salary"})
        return
    }

    // Update the expense amount
    expense.Amount = input.Amount

    if err := config.DB.Save(&expense).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update expense"})
        return
    }

    // Update the remaining salary by subtracting the expense difference
    salary.RemainingSalary -= expenseDifference
    if err := config.DB.Save(&salary).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update remaining salary"})
        return
    }

    // Return the updated expense and remaining salary
    c.JSON(http.StatusOK, gin.H{
        "updated expense":        expense.Amount,
		"updated expense name": expense.ExpenseName,
        "remaining salary":       salary.RemainingSalary,
    })
}

func DeleteExpense(c *gin.Context) {
    id := c.Param("id")
	
    var remove models.Expenses
    if err := config.DB.First(&remove, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
        return
    }

    if err := config.DB.Delete(&remove).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expense"})
        return
    }

    var Salary models.Salary
    if err := config.DB.First(&Salary).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Salary record not found"})
        return
    }

    Salary.RemainingSalary += remove.Amount

    if err := config.DB.Save(&Salary).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update remaining salary"})
        return
    }

    c.JSON(http.StatusAccepted, gin.H{
        "message": "Expense deleted",
        "expense id": remove.ID,
        "updated remaining salary": Salary.RemainingSalary,
    })
}

