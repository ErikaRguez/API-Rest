package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DNS = "host=172.17.0.2 user=postgres password=passwords dbname=universidad port=5432"


func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}
}
