package routes

import (
	"learning/todo/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteCategory(g *gin.RouterGroup, db *gorm.DB) {
	handler := handler.NewHandlerCategory(db)
	g.GET("/category", handler.GetCategory)
	g.GET("/category/:id", handler.GetCategoryByID)
	g.POST("/category", handler.SaveCategory)
	g.PUT("/category/:id", handler.UpdateCategory)
	g.DELETE("/category/:id", handler.DeleteCategory)

	return
}
