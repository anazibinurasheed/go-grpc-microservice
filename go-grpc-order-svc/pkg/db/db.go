package db

import (
	"log"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-order-svc/pkg/models"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Order{})
	return Handler{DB: db}

}
