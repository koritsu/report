package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"restapi-go/model/entities"
)

var db *gorm.DB

func InitDB() error {

	var err error
	db, err = gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	if !db.Migrator().HasTable(&entities.Content{}) {

		err := db.AutoMigrate(&entities.Content{})
		if err != nil {
			return err
		}
	}

	return nil
}

func DB() *gorm.DB {
	return db
}
