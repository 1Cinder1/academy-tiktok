package model

import "time"

type CommentSubject struct {
	Id             int64     `gorm:"column:comment_id" json:"comment_id" form:"comment_id" db:"comment_id"`
	CommentUserId  int64     `gorm:"column:comment_user_id" json:"comment_user_id" form:"comment_user_id" db:"comment_user_id"`
	CommentText    string    `gorm:"column:comment_text" json:"comment_text" form:"comment_text" db:"comment_text"`
	CommentVideoId int64     `gorm:"column:comment_video_id" json:"comment_video_id" form:"comment_video_id" db:"comment_video_id"`
	CommentStatus  int8      `gorm:"column:comment_status" json:"comment_status" form:"comment_status" db:"comment_status"`
	CreateTime     time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time" form:"update_time" db:"update_time"`
}
