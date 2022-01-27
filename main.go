package main

import (
	"github.com/Fadhli12/go-gin-gorm-playground/book"
	"github.com/Fadhli12/go-gin-gorm-playground/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	dbUser := envVariable("DB_USER")
	dbPassword := envVariable("DB_PASSWORD")
	dbHost := envVariable("DB_HOST")
	dbPort := envVariable("DB_PORT")
	dbTable := envVariable("DB_TABLE")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbTable + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}
	db.AutoMigrate(&book.Book{})

	//bookRepository := book.NewRepository(db)
	//bookService := book.NewService(bookRepository)

	router := gin.Default()
	handler.NewBookHandler(&handler.ConfigBook{router}, db)
	handler.NewAuthorHandler(&handler.ConfigAuthor{router}, db)
	router.Run()
}

func envVariable(key string) string {
	return os.Getenv(key)
}
