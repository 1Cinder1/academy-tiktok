package model

import "time"

type FollowSubject struct {
	Id           int       `gorm:"column:follow_id" json:"follow_id" form:"follow_id" db:"follow_id"`
	FollowerId   int       `gorm:"column:follower_id" json:"follower_id" form:"follower_id" db:"follower_id"`
	FollowingId  int       `gorm:"column:following_id" json:"following_id" form:"following_id" db:"following_id"`
	FollowStatus int8      `gorm:"column:follow_status" json:"follow_status" form:"follow_status" db:"follow_status"`
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time" form:"update_time" db:"update_time"`
}
