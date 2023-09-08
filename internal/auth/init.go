package auth

import (
	"crypto/rsa"
	"github.com/803Studio/kptl-user-center/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func Init() error {
	b, err := os.ReadFile(config.AppConfig.Keys.Private)
	if err != nil {
		return err
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return err
	}

	b, err = os.ReadFile(config.AppConfig.Keys.Public)
	if err != nil {
		return err
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return err
	}

	return nil
}
