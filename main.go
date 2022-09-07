package main

import (
	"fmt"
	"learning/todo/config"
	"learning/todo/routes"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.Connect(&gorm.DB{})

	fmt.Println(db)

	router := gin.Default()
	routes.Routes(router, db)
	router.Run("localhost:8080")
}
