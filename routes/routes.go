package routes

import (
	"github.com/gin-gonic/gin"
	"example/connecting/handler"
)

func Routes() *gin.Engine {
	router := gin.Default()

	router.POST("/add", handler.CreateExpense)
	router.POST("/createUser", handler.CreateUser)

	router.Run("localhost:8080")

	return router

}
