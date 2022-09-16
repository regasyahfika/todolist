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
	var request helper.PaginationRequest

	err := c.Bind(&request)

	if err != nil {
		response := helper.APIResponse("Failed get request todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todos, err := h.service.FindAll(request)

	if err != nil {
		response := helper.APIResponse("Failed get list todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get list todo", http.StatusOK, "success", todos)
	c.JSON(http.StatusOK, response)
	return
}

func (h *todoHandler) GetTodoByID(c *gin.Context) {
	var inputID todo.GetTodoID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed get detail todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todo, err := h.service.FindByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed get detail todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get detail todo", http.StatusOK, "success", todo)
	c.JSON(http.StatusOK, response)
	return
}

func (h *todoHandler) SaveTodo(c *gin.Context) {
	var input todo.Todo

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed create todo", http.StatusUnprocessableEntity, "error", errors)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	todo, err := h.service.Save(input)
	if err != nil {
		response := helper.APIResponse("Failed create todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create todo", http.StatusOK, "success", todo)
	c.JSON(http.StatusOK, response)
	return
}

func (h *todoHandler) UpdateTodo(c *gin.Context) {
	var inputID todo.GetTodoID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed get detail todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData todo.Todo
	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed update todo", http.StatusUnprocessableEntity, "error", errors)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateTodo, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed update todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success update todo", http.StatusOK, "success", updateTodo)
	c.JSON(http.StatusOK, response)
	return
}

func (h *todoHandler) DeleteTodo(c *gin.Context) {
	var inputID todo.GetTodoID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed get id todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleteTodo, err := h.service.Delete(inputID)
	if err != nil {
		response := helper.APIResponse("Failed delete todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete todo", http.StatusOK, "success", deleteTodo)
	c.JSON(http.StatusOK, response)
	return
}
