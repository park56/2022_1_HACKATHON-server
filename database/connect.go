package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {

	USER := "root"
	PASS := "12345678"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "2022_1_Hackerton"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("gorm DB Open Error")
		panic(err)
	}

	DB = db

	//DB.AutoMigrate(&models.User{})

}
