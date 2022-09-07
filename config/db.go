package config

import (
	"fmt"
	"learning/todo/category"
	"learning/todo/todo"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	config map[string]string
)

func Connect(db *gorm.DB) *gorm.DB {
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("Failed read .env")
	}

	credentials := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DATABASE_HOST"],
		config["DB_DATABASE"],
	)

	connect, err := gorm.Open(mysql.Open(credentials), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection error", err)
	}

	Migrate(connect)

	return connect

}

func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&category.Category{}, &todo.Todo{})

	return db
}
