package internal

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go_tiktok/app/global"
	"go_tiktok/app/internal/model"
	"gorm.io/gorm"
)

func GetMessageList(ctx context.Context, id int, user_id int) (error, []model.MessageSubject) {
	messages := []model.MessageSubject{}
	err := global.MysqlDB.
		WithContext(ctx).
		Table("message").
		Where("from_user_id = ? AND to_user_id = ?", id, user_id).
		Find(&messages).
		Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "follow_subject"),
			)
			return fmt.Errorf("internal err"), nil
		}
		return err, nil
	}
	return nil, messages
}
