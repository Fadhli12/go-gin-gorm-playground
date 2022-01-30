package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectToDb() (*gorm.DB, error) {
	dbUser := envVariable("DB_USER")
	dbPassword := envVariable("DB_PASSWORD")
	dbHost := envVariable("DB_HOST")
	dbPort := envVariable("DB_PORT")
	dbTable := envVariable("DB_TABLE")

	var connectionString = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost,
		dbPort, dbTable)
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func envVariable(key string) string {
	return os.Getenv(key)
}
