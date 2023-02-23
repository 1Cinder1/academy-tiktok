package internal

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go_tiktok/app/global"
	"go_tiktok/app/internal/model"
	"gorm.io/gorm"
)

func AddFollower(ctx context.Context, id int, to_user_id int) error {
	follow := model.FollowSubject{
		FollowerId:   id,
		FollowingId:  to_user_id,
		FollowStatus: 1,
	}
	err := global.MysqlDB.WithContext(ctx).Table("favourite").Create(&follow).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "video_subject"),
			)
			return fmt.Errorf("internal err")
		}
		return err
	}
	return nil
}

func DeleteFollower(ctx context.Context, id int, to_user_id int) error {
	follow := model.FollowSubject{
		FollowerId:  id,
		FollowingId: to_user_id,
	}
	err := global.MysqlDB.WithContext(ctx).Table("favourite").Delete(&follow).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "video_subject"),
			)
			return fmt.Errorf("internal err")
		}
		return err
	}
	return nil
}

func CheckIsFollower(ctx context.Context, id int, toUserId int) bool {
	follow := model.FollowSubject{
		FollowerId:  id,
		FollowingId: toUserId,
	}
	err := global.MysqlDB.WithContext(ctx).Table("favourite").Where(&follow).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "video_subject"),
			)
			return false
		}
		return true
	}
	return false
}
