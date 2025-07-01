package auth

import (
	"context"
	"errors"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

func (a *UseCase) ExtractClaims(_ context.Context, token *jwt.Token) (*entity.UserClaims, error) {
	claims, ok := token.Claims.(*entity.UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}
