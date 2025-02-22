package handler

import (
	"net/http"
	"nexmedis-technical-test/app/model/dto"
	"nexmedis-technical-test/app/usecase"

	"nexmedis-technical-test/app/helper"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication-related requests
type UserHandler struct {
	userUsecase usecase.UserUsecase
}

// NewAuthHandler creates a new AuthHandler with the provided usecase
func NewAuthHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) Login(c *gin.Context) {
	// Handle login request
	authCredential := dto.AuthRequest{}
	if err := c.ShouldBindJSON(&authCredential); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseBuilder(http.StatusBadRequest, err.Error(), "Invalid request payload"))
		return
	}

	if authCredential.Username == "" || authCredential.Password == "" {
		c.JSON(http.StatusBadRequest, helper.ResponseBuilder(http.StatusBadRequest, "username and password can be nil", "invalid request payload"))
		return
	}

	accessToken, refreshToken, err := h.userUsecase.Login(authCredential)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseBuilder(http.StatusInternalServerError, err.Error(), "failed to login"))
		return
	}

	tokens := dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, helper.ResponseBuilder(http.StatusOK, tokens, "Login success"))
}

func (h *UserHandler) Register(c *gin.Context) {
	// Handle register request
}
