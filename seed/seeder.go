package main

import (
	"bobee/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var users = []models.User{
	models.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Gender:   "male",
		Password: "10c4981bb793e1698a83aea43030a388",
		Photo:    "",
	},
}

// Conn ...
var Conn *gorm.DB

// main for exceute seed
func main() {
	conn := getDBConnection()
	defer conn.Close()
	Conn.LogMode(true)

	for _, k := range users {
		err := Conn.Table("users").Create(&k)
		if err != nil {
			log.Println("[Error] seeder.main : ", err)
		}
		err = nil
	}

}

func getDBConnection() *gorm.DB {
	conn, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	return conn
}

//Connect to db
func Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", "user=postgres password=Andrerose01 host=localhost dbname=bobee sslmode=disable")
	Conn = db
	return
}
