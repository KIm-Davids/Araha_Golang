package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func SubscriptionRepo() (*gorm.DB, error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/araha_Golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database")

	//if err := db.AutoMigrate(&models.Subscription{}); err != nil {
	//	log.Fatalf("Couldn't create a database : %v", err)
	//}
	return db, nil
}
