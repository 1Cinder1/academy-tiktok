package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go_tiktok/app/global"

	"time"
)

type CustomClaims struct {
	BufferTime int64
	jwt.RegisteredClaims
	BaseClaims
}

type BaseClaims struct {
	Id         int
	Username   string
	CreateTime time.Time
	UpdateTime time.Time
}

func GetClaims(secret string, token string) (*CustomClaims, error) {

	j := NewJWT(&Config{SecretKey: secret})
	claims, err := j.ParseToken(token)
	if err != nil {
		fmt.Print(err)
		err := errors.New("parse token failed")
		return nil, err
	}
	return claims, nil
}

func GetUserId(token string) (int, error) {
	jwtConfig := global.Config.Middleware.Jwt
	j := NewJWT(&Config{
		SecretKey: jwtConfig.SecretKey,
	})

	mc, err := j.ParseToken(token)
	if err != nil {
		return -1, err
	}
	return mc.Id, nil
}
