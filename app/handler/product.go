package handler

import (
	"net/http"
	"strconv"

	"nexmedis-technical-test/app/helper"
	"nexmedis-technical-test/app/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	// Get query parameters
	name := c.Query("name") // Search by name (optional)
	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1") // Default to page 1

	// Convert limit and page to integers
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, helper.ResponseBuilder(http.StatusBadRequest, err.Error(), "Invalid limit Parameter"))
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, helper.ResponseBuilder(http.StatusBadRequest, err.Error(), "Invalid Page Parameter"))
		return
	}

	// Calculate offset based on page number
	offset := (page - 1) * limit

	// Fetch products
	products, err := h.productUsecase.GetProducts(name, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseBuilder(http.StatusInternalServerError, err.Error(), "Failed to fetch products"))
		return
	}

	// Return response
	c.JSON(http.StatusOK, helper.ResponseBuilder(http.StatusOK, products, "Products retrieved successfully"))
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	// Get product ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, helper.ResponseBuilder(http.StatusBadRequest, err.Error(), "Invalid id"))
		return
	}

	// Fetch product by ID
	product, err := h.productUsecase.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseBuilder(http.StatusInternalServerError, err.Error(), "Failed to fetch product"))
		return
	}

	// Return response
	c.JSON(http.StatusOK, helper.ResponseBuilder(http.StatusOK, product, "Product retrieved successfully"))
}
