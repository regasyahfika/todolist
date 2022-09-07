package handler

import (
	"learning/todo/category"
	"learning/todo/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type categoryHandler struct {
	service category.Service
}

func NewHandlerCategory(db *gorm.DB) *categoryHandler {
	service := category.NewCategoryService(db)
	return &categoryHandler{service: service}
}

func (h *categoryHandler) GetCategory(c *gin.Context) {
	var request category.CategoryPaginationRequest

	err := c.Bind(&request)

	if err != nil {
		response := helper.APIResponse("Failed get list category", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	category, err := h.service.FindAll(request)
	if err != nil {
		response := helper.APIResponse("Failed get list category", http.StatusBadRequest, "error", err.Error())

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get list category", http.StatusOK, "success", category)
	c.JSON(http.StatusOK, response)
	return
}

func (h *categoryHandler) GetCategoryByID(c *gin.Context) {
	var input category.GetCategoryID

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed get detail category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	categoryDetail, err := h.service.FindByID(input)

	if err != nil {
		response := helper.APIResponse("Failed get detail category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get detail category", http.StatusOK, "success", categoryDetail)
	c.JSON(http.StatusOK, response)
	return
}

func (h *categoryHandler) SaveCategory(c *gin.Context) {
	var input category.Category

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed create category", http.StatusUnprocessableEntity, "error", errors)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	category, err := h.service.Save(input)
	if err != nil {
		response := helper.APIResponse("Failed create category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create category", http.StatusOK, "success", category)
	c.JSON(http.StatusOK, response)
	return

}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var inputID category.GetCategoryID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed get detail category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData category.Category

	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed update category", http.StatusUnprocessableEntity, "error", errors)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateCategory, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed get detail category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success update category", http.StatusOK, "success", updateCategory)
	c.JSON(http.StatusOK, response)
	return
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	var inputID category.GetCategoryID

	err := c.ShouldBindUri(inputID)
	if err != nil {
		response := helper.APIResponse("Failed delete category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updateCategory, err := h.service.Delete(inputID)

	response := helper.APIResponse("Success delete category", http.StatusOK, "success", updateCategory)
	c.JSON(http.StatusOK, response)
	return
}
