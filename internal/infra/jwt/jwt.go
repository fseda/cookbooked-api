package jwtutil

import (
	"time"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

// var secretKey = []byte("asdfasdfasdf")

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	ExpiresAt int64  `json:"exp,omitempty"`
}

func GenerateToken(user *models.User, secretKey []byte) (string, error) {
	claims := &CustomClaims{
		UserID:    user.ID,
		Username:  user.Username,
		Role:      string(user.Role),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ValidateToken(tokenStr string, secretKey []byte) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, globalerrors.AuthInvalidToken
	}

	return claims, nil
}
