package main

import (
	"example/connecting/config"
	"example/connecting/routes"
)


func main() {
	
	config.ConnectDB()
	routes.Routes()
	


	
}