package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(g *gin.Engine, db *gorm.DB) {
	api := g.Group("/api")
	RouteCategory(api, db)

	return
}
