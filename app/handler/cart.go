package handler

import (
	"net/http"
	"nexmedis-technical-test/app/helper"
	"nexmedis-technical-test/app/model/dto"
	"nexmedis-technical-test/app/usecase"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUsecase usecase.CartUsecase
}

func NewCartHandler(cartUsecase usecase.CartUsecase) *CartHandler {
	return &CartHandler{cartUsecase: cartUsecase}
}

// Add to Cart Handler (POST with JSON Body)
func (h *CartHandler) AddToCart(c *gin.Context) {
	// Extract user ID from token (assuming middleware sets it in context)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, helper.ResponseBuilder(http.StatusUnauthorized, "", "Unauthorized"))
		return
	}

	var req dto.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseBuilder(http.StatusBadRequest, err.Error(), "Bad request payload"))
		return
	}

	err := h.cartUsecase.AddToCart(userID.(int), req.ProductID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseBuilder(http.StatusInternalServerError, err.Error(), "error occurs"))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseBuilder(http.StatusOK, "", "Success add to cart"))
}

// Get Cart Items Handler (GET)
func (h *CartHandler) GetCartItems(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, helper.ResponseBuilder(http.StatusUnauthorized, "", "Unauthorized"))
		return
	}

	cartItems, err := h.cartUsecase.GetCartItems(userID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, helper.ResponseBuilder(http.StatusNotFound, err.Error(), "No Item Found"))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseBuilder(http.StatusOK, cartItems, "success get cart"))
}

func (h *CartHandler) Checkout(c *gin.Context) {
	userID := c.MustGet("user_id").(int) // Assuming user_id is extracted from JWT token

	orderID, err := h.cartUsecase.Checkout(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseBuilder(http.StatusInternalServerError, err.Error(), "error occurs"))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseBuilder(http.StatusOK, orderID, "Checkout successful"))
}
