package routes

import (
	"github.com/gin-gonic/gin"
	"example/connecting/handler"
)

func Routes() *gin.Engine {
	router := gin.Default()

	router.POST("/add", handler.CreateExpense)
	router.GET("/all", handler.GetAllExpense)
	router.GET("/expense/:id", handler.GetExpenseByID)
	router.PUT("/update/:id", handler.UpdateAmount)
	router.DELETE("delete/:id",handler.DeleteExpense)

	// user 
	router.POST("/salary", handler.UserSalary)

	router.Run("localhost:8080")

	return router

}
