package db

import (
	"github.com/iamtakdir/url-shortner-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	//changed database config
	dsn := "host=localhost user=postgres password=843543 dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database Connected Successfully")
	DB = db
	db.AutoMigrate(&models.ShortUrlModel{})
}
