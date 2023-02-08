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
	Id            int
	Name          string
	FollowCount   int64
	FollowerCoune int64
	IsFollow      bool
}

func Getuser(ctx context.Context, userId int, id int) (error, *User) {
	user := &model.UserSubject{}
	err := global.MysqlDB.WithContext(ctx).
		Table("user").
		Where("user_id = ?", userId).
		First(&user).Error

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user_subject"),
			)
			return fmt.Errorf("internal err"), nil
		} else {
			return fmt.Errorf("user doesn't exist"), nil
		}
	}

	err1, followCount := GetFollowCount(ctx, userId)
	if err1 != nil {
		if err1 != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err1),
				zap.String("table", "follow_subject"),
			)
			return fmt.Errorf("internal err"), nil
		}

	}

	err2, followerCount := GetFollowerCount(ctx, userId)
	if err2 != nil {
		if err2 != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err2),
				zap.String("table", "follow_subject"),
			)
			return fmt.Errorf("internal err"), nil
		}

	}

	err3, flag := CheckIsFollow(ctx, id, userId)
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
	user1 := User{
		Id:            user.Id,
		Name:          user.Username,
		FollowCount:   followCount,
		FollowerCoune: followerCount,
		IsFollow:      flag,
	}
	return nil, &user1
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
