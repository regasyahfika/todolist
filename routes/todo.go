package routes

import (
	"learning/todo/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteTodo(g *gin.RouterGroup, db *gorm.DB) {
	handler := handler.NewHandlerTodo(db)

	g.GET("/todo", handler.GetTodo)

	return
}
