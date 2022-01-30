package main

import (
	"fmt"
	"github.com/Fadhli12/go-gin-gorm-playground/app/handler"
	"github.com/Fadhli12/go-gin-gorm-playground/config"
	"github.com/Fadhli12/go-gin-gorm-playground/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
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
	handler.NewLoginHandler(&handler.ConfigLogin{router}, db)
	handler.NewBookHandler(&handler.ConfigBook{router}, db)
	handler.NewAuthorHandler(&handler.ConfigAuthor{router}, db)

	port := os.Getenv("WEB_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
