package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := "host=localhost user=postgres password=Tujhcerf223 dbname=jwt port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ")
		log.Fatal("connection error:", err)
	}

	DB.AutoMigrate(&User{})

}

//dsn := "host=localhost user=postgres password=Tujhcerf223 dbname=test port=5432 sslmode=disable"
//DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
