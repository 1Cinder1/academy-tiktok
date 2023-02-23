package internal

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go_tiktok/app/global"
	"go_tiktok/app/internal/model"
	"gorm.io/gorm"
)

type VideoList struct {
	Id             int
	Author         *Author
	PlayUrl        string
	CoverUrl       string
	FavouriteCount int64
	CommentCount   int64
	IsFavourite    bool
	Title          string
}
type Author struct {
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

func GetVideoList(ctx context.Context, id int, user_id int) (error, []VideoList) {
	videos := []model.VideoSubject{}
	VideoLists := []VideoList{}
	err := global.MysqlDB.WithContext(ctx).
		Table("video").
		Where("author_id = ?", user_id).
		Find(videos).
		Error

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			global.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "video_subject"),
			)
			return fmt.Errorf("internal err"), nil
		}
		return err, nil
	}

	for _, v := range videos {
		user := &model.UserSubject{}
		err0 := global.MysqlDB.WithContext(ctx).
			Table("user").
			Where("user_id = ?", user_id).
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

		err1, followCount := GetFollowCount(ctx, user_id)
		if err1 != nil {
			if err1 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err1),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err2, followerCount := GetFollowerCount(ctx, user_id)
		if err2 != nil {
			if err2 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err2),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err3, flag := CheckIsFollow(ctx, id, user_id)
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

		err4, favouritecount := GetFavouriteCount(ctx, user_id)
		if err4 != nil {
			if err4 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err4),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err5, workcount := GetWorkCount(ctx, user_id)
		if err5 != nil {
			if err5 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err5),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err6, totalfavouritecount := GetTotalFavouritedCount(ctx, user_id)
		if err6 != nil {
			if err6 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err6),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err7, videoFavouriteCount := GetVideoFavouriteCount(ctx, v.Id)
		if err7 != nil {
			if err7 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err7),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}

		err8, videoCommentCount := GetVideoComentCount(ctx, v.Id)
		if err8 != nil {
			if err8 != gorm.ErrRecordNotFound {
				global.Logger.Error("query mysql record failed.",
					zap.Error(err8),
					zap.String("table", "follow_subject"),
				)
				return fmt.Errorf("internal err"), nil
			}

		}
		author := &Author{
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

		videoList := VideoList{
			Id:             v.Id,
			Author:         author,
			PlayUrl:        v.PlayUrl,
			CoverUrl:       v.CoverUrl,
			FavouriteCount: videoFavouriteCount,
			CommentCount:   videoCommentCount,
			IsFavourite:    false,
			Title:          v.Title,
		}
		VideoLists = append(VideoLists, videoList)
	}

	return nil, VideoLists
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

func GetVideoFavouriteCount(ctx context.Context, videoId int) (error, int64) {
	var count int64
	err := global.MysqlDB.WithContext(ctx).
		Table("favourite").
		Where("video_id = ?", videoId).
		Count(&count).Error

	if err != nil {
		return err, 0
	}
	return nil, count
}

func GetVideoComentCount(ctx context.Context, videoId int) (error, int64) {
	var count int64
	err := global.MysqlDB.WithContext(ctx).
		Table("comment").
		Where("comment_video_id = ?", videoId).
		Count(&count).Error

	if err != nil {
		return err, 0
	}
	return nil, count
}
