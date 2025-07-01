package auth

import (
	"context"
	"fmt"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

func (a *UseCase) ParseToken(_ context.Context, tokenStr string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenStr, &entity.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.secretKey, nil
	})
}
