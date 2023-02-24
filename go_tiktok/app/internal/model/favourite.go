package model

import "time"

type FavouriteSubject struct {
	Id              int64     `gorm:"column:favourite_id" json:"favourite_id" form:"favourite_id" db:"favourite_id"`
	UserId          int       `gorm:"column:user_id" json:"user_id" form:"user_id" db:"user_id"`
	FavouritUserId  int       `gorm:"column:favourite_user_id" json:"favourite_user_id" form:"favourite_user_id" db:"favourite_user_id"`
	VideoId         int       `gorm:"column:video_id" json:"video_id" form:"video_id" db:"video_id"`
	FavouriteStatus int8      `gorm:"column:favourite_status" json:"favourite_status" form:"favourite_status" db:"favourite_status"`
	CreateTime      time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time" form:"update_time" db:"update_time"`
}
