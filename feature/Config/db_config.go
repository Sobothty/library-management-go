package Config

import (
	"Go-Project/feature/Model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dns := "host=localhost user=postgres password=qwer dbname=go-project port=5432"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	error := db.AutoMigrate(&Model.Book{})
	if error != nil {
		return nil
	}
	return db
}
