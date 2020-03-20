package utils

import (
	"log"

	"github.com/jinzhu/gorm"
)

func GetDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}
