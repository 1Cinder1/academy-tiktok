package model

import "time"

type VideoSubject struct {
	Id         int       `gorm:"column:video_id" json:"video_id" form:"video_id" db:"video_id"`
	AuthorId   int       `gorm:"column:author_id" json:"author_id" form:"author_id" db:"author_id"`
	CoverUrl   string    `gorm:"column:cover_url" json:"cover_url" form:"cover_url" db:"cover_url"`
	PlayUrl    string    `gorm:"column:play_url" json:"play_url" form:"play_url" db:"play_url"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time" form:"update_time" db:"update_time"`
}
