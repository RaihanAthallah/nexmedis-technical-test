package usecase

import (
	"fmt"
	"nexmedis-technical-test/app/auth"
	"nexmedis-technical-test/app/helper"
	"nexmedis-technical-test/app/model/dto"
	repository "nexmedis-technical-test/app/repository"
	"strconv"
	"sync"
)

type UserUsecase interface {
	Login(authCredential dto.AuthRequest) (string, string, error)
	Register(username string, email string, password string) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) Login(authCredential dto.AuthRequest) (string, string, error) {
	// Retrieve user from the repository
	user, err := u.userRepo.GetUserByUsername(authCredential.Username)
	if err != nil {
		return "", "", err // User not found
	}

	// Verify password
	if err := helper.VerifyPassword(user.Password, authCredential.Password); err != nil {
		return "", "", err // Password does not match
	}

	// userid from int to string
	userID := strconv.Itoa(user.ID)
	encryptedUserID, err := helper.EncryptString(userID)
	if err != nil {
		return "", "", err
	}
	var accessToken, refreshToken string
	var wg sync.WaitGroup
	var tokenErr error

	wg.Add(2)
	go func() {
		fmt.Println("encrypting access token", encryptedUserID)
		defer wg.Done()
		accessToken, tokenErr = auth.GenerateAccessToken(user.Username, user.Email, encryptedUserID)
	}()
	go func() {
		defer wg.Done()
		refreshToken, tokenErr = auth.GenerateRefreshToken(user.Username, user.Email, encryptedUserID)
	}()

	wg.Wait()

	if tokenErr != nil {
		return "", "", tokenErr
	}

	return accessToken, refreshToken, nil
}

func (u *userUsecase) Register(username string, email string, password string) error {
	return nil
}
