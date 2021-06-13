package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectionDB() {
	db, err := gorm.Open("sqlite3", "./gorm.db")

	if err != nil {
		panic("Failed to connect to DB")
	}

	// defer db.Close()

	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Permission{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&Experience{})
	db.AutoMigrate(&Technology{})

	DB = db
}
