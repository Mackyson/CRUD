package storage

import (
	"log"

	"CRUD/model"
	"CRUD/pkg/DButil"
)

func Migrate() {
	db := DButil.GetClient()
	if !db.HasTable(&model.User{}) {
		err := db.CreateTable(&model.User{})
		if err != nil {
			log.Printf("User table exists")
		}
	}
	db.LogMode(true)
	db.AutoMigrate(&model.User{})
}
