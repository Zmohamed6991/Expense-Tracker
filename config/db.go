//where to connect the database

package config

import (
	"example/connecting/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(`postgres://postgres:password@localhost:5432/Expense Tracker`)) //username, password, db name
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Expenses{}, &models.Categories{})
	
	
	DB = db


}


