package internal

import (
	"context"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	g "go_tiktok/app/global"
	"go_tiktok/app/internal/model"
	"golang.org/x/crypto/sha3"
	"gorm.io/gorm"
)

func CheckUserIsExist(ctx context.Context, username string) error {
	userSubject := &model.UserSubject{}
	err := g.MysqlDB.WithContext(ctx).
		Table("user").
		Select("username").
		Where("username = ?", username).
		First(userSubject).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user_subject"),
			)
			return fmt.Errorf("internal err")
		}
	} else {
		return fmt.Errorf("username already exist")
	}

	return nil
}

func EncryptPassword(password string) string {
	d := sha3.Sum224([]byte(password))
	return hex.EncodeToString(d[:])
}

func CreateUser(ctx context.Context, userSubject *model.UserSubject) {
	g.MysqlDB.WithContext(ctx).
		Table("user").
		Create(userSubject)
}

func GeeUserId(ctx context.Context, username string) int {
	user := model.UserSubject{}
	g.MysqlDB.WithContext(ctx).
		Table("user").
		Where("username = ?", username).
		First(&user)
	return user.Id
}
