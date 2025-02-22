package handler

import (
	"net/http"
	"nexmedis-technical-test/app/helper"
	"nexmedis-technical-test/app/model/dto"
	"nexmedis-technical-test/app/usecase"

	"github.com/gin-gonic/gin"
)

type BankHandler struct {
	bankUsecase usecase.BankUsecase
}

func NewBankHandler(bankUsecase usecase.BankUsecase) *BankHandler {
	return &BankHandler{bankUsecase: bankUsecase}
}

func (h *BankHandler) Deposit(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, helper.ResponseBuilder(http.StatusUnauthorized, "", "Unauthorized"))
		return
	}

	// Parse request body
	var request dto.TransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseBuilder(http.StatusBadRequest, err.Error(), "Bad request payload"))
		return
	}

	// Perform deposit
	err := h.bankUsecase.Deposit(userID.(int), request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseBuilder(http.StatusInternalServerError, err.Error(), "error occurs"))
		return
	}

	// Success response
	c.JSON(http.StatusOK, helper.ResponseBuilder(http.StatusOK, "", "Deposit successful"))
}

func (h *BankHandler) Withdraw(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, helper.ResponseBuilder(http.StatusUnauthorized, "", "Unauthorized"))
		return
	}

	// Parse request body
	var request dto.TransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseBuilder(http.StatusBadRequest, err.Error(), "Bad request payload"))
		return
	}

	// Perform deposit
	err := h.bankUsecase.Withdraw(userID.(int), request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseBuilder(http.StatusInternalServerError, err.Error(), "error occurs"))
		return
	}

	// Success response
	c.JSON(http.StatusOK, helper.ResponseBuilder(http.StatusOK, "", "Withdraw successful"))
}

func (h *BankHandler) GetBalance(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, helper.ResponseBuilder(http.StatusUnauthorized, "", "Unauthorized"))
		return
	}

	balance, err := h.bankUsecase.GetDetails(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseBuilder(http.StatusInternalServerError, err.Error(), "error occurs"))
		return
	}

	// Success response
	c.JSON(http.StatusOK, helper.ResponseBuilder(http.StatusOK, balance, "Get Accound Balance Successfull"))
}
