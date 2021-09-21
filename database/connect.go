package database

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Conn ...
var Conn *gorm.DB

// Connect ...
func Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", beego.AppConfig.String("sqlconn"))
	Conn = db
	return
}
