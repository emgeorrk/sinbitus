package auth

import (
	"context"

	"github.com/emgeorrk/sinbitus/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

func (a *UseCase) GenerateToken(_ context.Context, userID uint64, username string) (string, error) {
	issTS := a.time.Now()
	expTS := issTS.Add(a.ttl)

	claims := entity.UserClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTS),
			IssuedAt:  jwt.NewNumericDate(issTS),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secretKey)
}
