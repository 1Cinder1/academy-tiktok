package internal

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go_tiktok/app/global"
	"go_tiktok/app/internal/model"
	"gorm.io/gorm"
)

func SendMessage(ctx context.Context, id int, user_id int, message string) error {
	message1 := model.MessageSubject{
		FromUserId: id,
		Message:    message,
		ToUserId:   user_id,
	}

	err := global.MysqlDB.WithContext(ctx).Table("message").Create(&message1).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "follow_subject"),
			)
			return fmt.Errorf("internal err")
		}
		return err
	}
	return nil
}
