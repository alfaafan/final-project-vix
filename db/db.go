package db

import (
	"finalProject/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error


func DatabaseInit() {
	host := config.GetConfig().DB_HOST
	user := config.GetConfig().DB_USERNAME
	password := config.GetConfig().DB_PASSWORD
	dbName := config.GetConfig().DB_NAME
	port := config.GetConfig().DB_PORT

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}

