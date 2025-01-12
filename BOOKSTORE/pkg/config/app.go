package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import the MySQL dialect
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "med:med99@/bookdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
