package main

import (
	"finalProject/db"
	"finalProject/routes"
)

func main() {
	db.DatabaseInit()
	gorm := db.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()


	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
