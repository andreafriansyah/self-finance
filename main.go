package main

import (
	"bobee/database"
	"bobee/models"
	_ "bobee/routers"
	"encoding/gob"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

func init() {
	gob.Register(models.User{})
}

func main() {
	conn := getDBConnection()
	defer conn.Close()
	database.Conn.LogMode(true)
	beego.Run()
}

func getDBConnection() *gorm.DB {
	conn, err := database.Connect()
	if err != nil {
		panic(err.Error())
	}
	return conn
}
