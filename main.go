package main

import (
	"fmt"
	"github.com/Fadhli12/go-gin-gorm-playground/app/handler"
	"github.com/Fadhli12/go-gin-gorm-playground/config"
	"github.com/Fadhli12/go-gin-gorm-playground/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()
	db, err := config.ConnectToDb()
	if err != nil {
		log.Fatal("DB connection error")
	}
	for _, model := range model.RegisterModels() {
		err = db.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("DB Migration successfully")

	router := gin.Default()
	handler.NewBookHandler(&handler.ConfigBook{router}, db)
	handler.NewAuthorHandler(&handler.ConfigAuthor{router}, db)
	router.Run(":8080")
}
