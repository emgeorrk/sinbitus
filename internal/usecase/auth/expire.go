package auth

import (
	"context"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

func (a *UseCase) IsExpired(_ context.Context, token *jwt.Token) bool {
	nowTS := a.time.Now()

	claims, ok := token.Claims.(*entity.UserClaims)
	if !ok {
		return true
	}

	return claims.ExpiresAt != nil && claims.ExpiresAt.Before(nowTS)
}
