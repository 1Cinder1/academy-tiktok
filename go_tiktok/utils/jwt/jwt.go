package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWT struct {
	Config *Config
}

type Config struct {
	SecretKey   string
	ExpiresTime int64
	BufferTime  int64
	Issuer      string
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT(config *Config) *JWT {
	return &JWT{Config: config}
}
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {

	signingKey := []byte(j.Config.SecretKey)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return signingKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}

		}
	}
	if token != nil {
		if claim, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claim, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

func (j *JWT) CreatClaims(baseClaims *BaseClaims) CustomClaims {
	claims := CustomClaims{
		BufferTime: j.Config.BufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Truncate(time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.Config.ExpiresTime) * time.Second)),
			Issuer:    j.Config.Issuer,
		},
		BaseClaims: *baseClaims,
	}
	return claims
}

func (j *JWT) GenerateToken(claims *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	signingKey := []byte(j.Config.SecretKey)
	return token.SignedString(signingKey)
}
