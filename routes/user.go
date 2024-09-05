package routes

import (
	"example/connecting/handler"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.GET("/find", handler.GetExpense)
	router.POST("/addexpense", handler.CreateExpense)
	router.PUT("/:id/update",)
	router.DELETE("/id/delete",)
	
}
