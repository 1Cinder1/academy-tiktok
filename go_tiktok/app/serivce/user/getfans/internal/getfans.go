package internal

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go_tiktok/app/global"
	"go_tiktok/app/internal/model"
	"gorm.io/gorm"
)

type User struct {
	Id            int    `gorm:"column:user_id" json:"user_id" form:"user_id" db:"user_id"`
	Name          string `gorm:"column:username" json:"username" form:"username" db:"username"`
	FollowCount   int64
	FollowerCoune int64
	IsFollow      bool
}

func GetFans(ctx context.Context, userId int, id int) (error, []User) {
	follow := []model.FollowSubject{}
	user1 := []User{}
	err := global.MysqlDB.WithContext(ctx).
		Table("follow").
		Where("following_id = ?", userId).
		Find(&follow).Error
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

	for _, v := range follow {
		user := &model.UserSubject{}
		err0 := global.MysqlDB.WithContext(ctx).
			Table("user").
			Where("user_id = ?", v.FollowerId).
			First(&user).Error

		if err0 != nil {
			if err0 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err),
					zap.String("table", "user_subject"),
				)
				return fmt.Errorf("internal err"), nil
			} else {
				return err0, nil
			}
		}

		err1, followCount := GetFollowCount(ctx, v.FollowerId)
		if err1 != nil {
			if err1 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err1),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err2, followerCount := GetFollowerCount(ctx, v.FollowerId)
		if err2 != nil {
			if err2 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err2),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err3, flag := CheckIsFollow(ctx, id, v.FollowerId)
		if err3 != nil {
			if err3 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err3),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}
			return err3, nil
		}
		users := User{
			Id:            v.FollowerId,
			Name:          user.Username,
			FollowCount:   followCount,
			FollowerCoune: followerCount,
			IsFollow:      flag,
		}
		user1 = append(user1, users)
	}

	return nil, user1
}

func GetFollowCount(ctx context.Context, userId int) (error, int64) {
	var count int64
	err := global.MysqlDB.WithContext(ctx).
		Table("follow").
		Where("follower_id = ? AND follow_status = ?", userId, 1).
		Count(&count).Error

	if err != nil {
		return err, 0
	}
	return nil, count
}

func GetFollowerCount(ctx context.Context, userId int) (error, int64) {
	var count int64
	err := global.MysqlDB.WithContext(ctx).
		Table("follow").
		Where("following_id = ? AND follow_status = ?", userId, 1).
		Count(&count).Error

	if err != nil {
		return err, 0
	}
	return nil, count
}

func CheckIsFollow(ctx context.Context, followerId int, followingId int) (error, bool) {
	var count int64
	err := global.MysqlDB.WithContext(ctx).
		Table("follow").
		Where("following_id = ? AND follower_id = ?", followingId, followerId).
		Count(&count).Error

	if err != nil {
		return err, false
	}
	if count == 0 {
		return nil, false
	}
	return nil, true
}
