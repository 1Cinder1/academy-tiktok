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
	UserId          int
	Username        string
	FollowCount     int64
	FollowerCount   int64
	IsFollow        bool
	Avatar          string
	BackGroundImage string
	Signature       string
	TotalFavourited string
	WorkCount       int64
	FavouriteCount  int64
}

func GtFollowList(ctx context.Context, id int, user_id int) (error, []User) {
	follow := []model.FollowSubject{}
	followlist := []User{}
	err := global.MysqlDB.WithContext(ctx).
		Table("follow").
		Where("follower_id = ?", user_id).
		Find(follow).
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

	for _, v := range follow {
		user := &model.UserSubject{}
		err0 := global.MysqlDB.WithContext(ctx).
			Table("user").
			Where("user_id = ?", v.FollowingId).
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

		err1, followCount := GetFollowCount(ctx, v.FollowingId)
		if err1 != nil {
			if err1 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err1),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err2, followerCount := GetFollowerCount(ctx, v.FollowingId)
		if err2 != nil {
			if err2 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err2),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err3, flag := CheckIsFollow(ctx, id, v.FollowingId)
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

		err4, favouritecount := GetFavouriteCount(ctx, v.FollowingId)
		if err4 != nil {
			if err4 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err4),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err5, workcount := GetWorkCount(ctx, v.FollowingId)
		if err5 != nil {
			if err5 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err5),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err6, totalfavouritecount := GetTotalFavouritedCount(ctx, v.FollowingId)
		if err6 != nil {
			if err6 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err6),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		follow := User{
			UserId:          user.Id,
			Username:        user.Username,
			FollowCount:     followCount,
			FollowerCount:   followerCount,
			IsFollow:        flag,
			Avatar:          user.Avatar,
			BackGroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavourited: string(totalfavouritecount),
			WorkCount:       workcount,
			FavouriteCount:  favouritecount,
		}
		followlist = append(followlist, follow)
	}

	return nil, followlist
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

func GetWorkCount(ctx context.Context, userId int) (error, int64) {
	var count int64
	err := global.MysqlDB.WithContext(ctx).
		Table("video").
		Where("author_id = ?", userId).
		Count(&count).Error

	if err != nil {
		return err, 0
	}
	return nil, count
}

func GetFavouriteCount(ctx context.Context, userId int) (error, int64) {
	var count int64
	err := global.MysqlDB.WithContext(ctx).
		Table("favourite").
		Where("user_id = ?", userId).
		Count(&count).Error

	if err != nil {
		return err, 0
	}
	return nil, count
}

func GetTotalFavouritedCount(ctx context.Context, userId int) (error, int64) {
	var count int64
	err := global.MysqlDB.WithContext(ctx).
		Table("favourite").
		Where("favourite_user_id = ?", userId).
		Count(&count).Error

	if err != nil {
		return err, 0
	}
	return nil, count
}
