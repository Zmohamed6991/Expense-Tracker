package main

import (
	"example/connecting/config"
	

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	config.ConnectDB()

	r.Run(":8080") 
}