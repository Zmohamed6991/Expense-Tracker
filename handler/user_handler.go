package handler

import (
	"net/http"
	"example/connecting/models"
	"github.com/gin-gonic/gin"
	"example/connecting/config"

	
)

func GetExpense(c *gin.Context) {



	
}

func CreateExpense(c *gin.Context){

	var AddExpense *models.Expenses
	
	if err := c.BindJSON(&AddExpense); err != nil {
		c.JSON(http.StatusNotFound, AddExpense)
		return
	}
	
	c.JSON(http.StatusCreated, AddExpense)

	err := config.DB.Create(&AddExpense) 
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	



	
}

	

func UpdateExpense(c *gin.Context) {

}

func DeleteExpense(c *gin.Context){

}