package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	accessTokenSecret  = []byte("your_access_token_secret")  // Change to a secure secret
	refreshTokenSecret = []byte("your_refresh_token_secret") // Change to a secure secret
)

// JWTClaims defines the structure of JWT claims
type JWTClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       string `json:"id"`
	jwt.StandardClaims
}

// NewJWTClaims creates a new JWTClaims instance
func NewJWTClaims(username, email, id string, expireDuration time.Duration) *JWTClaims {
	fmt.Println("id string", id)
	return &JWTClaims{
		Username: username,
		Email:    email,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
}

// GenerateAccessToken creates a short-lived JWT access token
func GenerateAccessToken(username, email, id string) (string, error) {
	claims := NewJWTClaims(username, email, id, 15*time.Minute) // 15 minutes validity
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessTokenSecret)
}

// GenerateRefreshToken creates a long-lived JWT refresh token
func GenerateRefreshToken(username, email, id string) (string, error) {
	claims := NewJWTClaims(username, email, id, 7*24*time.Hour) // 7 days validity
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshTokenSecret)
}

// ValidateToken parses and validates a JWT token
func ValidateToken(tokenString string, secret []byte) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ValidateAccessToken validates an access token
func ValidateAccessToken(tokenString string) (*JWTClaims, error) {
	return ValidateToken(tokenString, accessTokenSecret)
}

// ValidateRefreshToken validates a refresh token
func ValidateRefreshToken(tokenString string) (*JWTClaims, error) {
	return ValidateToken(tokenString, refreshTokenSecret)
}

// RefreshAccessToken issues a new access token using a valid refresh token
func RefreshAccessToken(refreshToken string) (string, error) {
	// Validate the refresh token
	claims, err := ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", errors.New("invalid or expired refresh token")
	}

	// Generate a new access token
	newAccessToken, err := GenerateAccessToken(claims.Username, claims.Email, claims.ID)
	if err != nil {
		return "", errors.New("failed to generate new access token")
	}

	return newAccessToken, nil
}
