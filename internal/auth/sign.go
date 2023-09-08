package auth

import (
	"fmt"
	"github.com/803Studio/kptl-user-center/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func Sign(info *model.UserAccount, keep bool) (string, error) {
	exp := time.Hour * 24
	if keep {
		exp = time.Hour * 24 * 7
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "cdl",
		Subject:   fmt.Sprintf("%d", info.Role),
		Audience:  []string{""},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        fmt.Sprintf("%d", info.Id),
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
}
