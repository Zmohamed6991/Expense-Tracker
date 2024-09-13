package handler

import (
	"example/connecting/config"
	"example/connecting/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllExpense(c *gin.Context) {
	var expenses []models.Expenses

	err := config.DB.Find(&expenses)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(200, expenses)
}

func GetExpenseByID(c *gin.Context){
	id := c.Param("id")

	err := c.BindJSON(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	var expense models.Expenses
	if err := config.DB.First(&expense, id); err != nil{
		c.JSON(http.StatusNotFound, err)
	}
	
}

func CreateExpense(c *gin.Context){
	var AddExpense models.Expenses
	
	err := c.BindJSON(&AddExpense)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": "invalid data"})
		return
	}
		
	if err := config.DB.Create(&AddExpense); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusCreated, "Success")
}

func CreateUser(c *gin.Context){
	var newUser []models.User

	err := c.BindJSON(&newUser)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
	}

	config.DB.Create(&newUser)
	if err != nil{
		c.JSON(http.StatusBadRequest,err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success"})
}

	

func UpdateAmount(c *gin.Context) {
	// get id
	id := c.Param("id")

	//get data off req body
	err := c.BindJSON(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	//find post were updating 
	var expense models.Expenses
	if err := config.DB.Find(&expense, id); err != nil{
		c.JSON(http.StatusBadRequest, err)
	}
	
	// update it
	if err := config.DB.Model(&expense).Updates(models.Expenses{
		Amount: expense.Amount,
	}); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
}

func DeleteExpense(c *gin.Context){

}