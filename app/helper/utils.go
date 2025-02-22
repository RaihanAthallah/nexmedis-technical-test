package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"nexmedis-technical-test/app/model/dto"
)

func HashPassword(password string) (string, error) {
	return "", nil
}

func VerifyPassword(hashedPassword, password string) error {
	return nil
}

func ResponseBuilder(status int, data interface{}, message string) *dto.GlobalResponse {
	return &dto.GlobalResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

var key = []byte("your-32-byte-secret-key-here!212") // AES-256 requires a 32-byte key

func EncryptString(str string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(str), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptString(str string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("invalid ciphertext length")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
