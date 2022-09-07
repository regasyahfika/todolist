package handler

import (
	"learning/todo/helper"
	"learning/todo/todo"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type todoHandler struct {
	service todo.Service
}

func NewHandlerTodo(db *gorm.DB) *todoHandler {
	service := todo.NewService(db)
	return &todoHandler{service: service}
}

func (h *todoHandler) GetTodo(c *gin.Context) {
	todo, err := h.service.FindAll()

	if err != nil {
		response := helper.APIResponse("Failed get list todo", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get list todo", http.StatusOK, "success", todo)
	c.JSON(http.StatusOK, response)
	return
}
