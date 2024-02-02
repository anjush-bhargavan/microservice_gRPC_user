package db

import (
	"log"
	"os"

	"github.com/anjush-bhargavan/microservice_gRPC_user/pkg/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn,ok := os.LookupEnv("DB_Config")
	if !ok {
		log.Fatal("error loading database env")
	}

	DB, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatalf("error conncting to databse: %v",err.Error())
	}
	DB.AutoMigrate(&entities.User{})

	return DB
}	