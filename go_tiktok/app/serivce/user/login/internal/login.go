package internal

import (
	"context"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	g "go_tiktok/app/global"
	"go_tiktok/app/internal/model"
	"go_tiktok/utils/jwt"
	"golang.org/x/crypto/sha3"
	"gorm.io/gorm"
	"time"
)

func CheckPassword(ctx context.Context, userSubject *model.UserSubject) error {
	err := g.MysqlDB.WithContext(ctx).
		Table("user").
		Where(&model.UserSubject{
			Username: userSubject.Username,
			Password: userSubject.Password,
		}).
		First(userSubject).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user"),
			)
			return fmt.Errorf("internal err")
		} else {
			return fmt.Errorf("invalid username or password")
		}
	}

	return nil
}

func GenerateToken(ctx context.Context, userSubject *model.UserSubject) (string, error) {
	jwtConfig := g.Config.Middleware.Jwt

	j := jwt.NewJWT(&jwt.Config{
		SecretKey:   jwtConfig.SecretKey,
		ExpiresTime: jwtConfig.ExpiresTime,
		BufferTime:  jwtConfig.BufferTime,
		Issuer:      jwtConfig.Issuer})
	claims := j.CreatClaims(&jwt.BaseClaims{
		Id:         userSubject.Id,
		Username:   userSubject.Username,
		CreateTime: userSubject.CreateTime,
		UpdateTime: userSubject.UpdateTime,
	})

	tokenString, err := j.GenerateToken(&claims)
	if err != nil {
		g.Logger.Error("generate token failed.", zap.Error(err))
		return "", fmt.Errorf("internal err")
	}

	err = g.Rdb.Set(ctx,
		fmt.Sprintf("jwt:%d", userSubject.Id),
		tokenString,
		time.Duration(jwtConfig.ExpiresTime)*time.Second).Err()
	if err != nil {
		g.Logger.Error("set redis cache failed.",
			zap.Error(err),
			zap.String("key", "jwt:[id]"),
			zap.Int("id", userSubject.Id),
		)
		return "", fmt.Errorf("internal err")
	}

	return tokenString, nil
}

func EncryptPassword(password string) string {
	d := sha3.Sum224([]byte(password))
	return hex.EncodeToString(d[:])
}

func GeeUserId(ctx context.Context, username string) int {
	user := model.UserSubject{}
	g.MysqlDB.WithContext(ctx).
		Table("user").
		Where("username = ?", username).
		First(&user)
	return user.Id
}
