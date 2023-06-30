package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sebastianpacuk/taskrocket/taskrocket-tasks/pkg/models"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal(err)
	}

	return Handler{db}
}
