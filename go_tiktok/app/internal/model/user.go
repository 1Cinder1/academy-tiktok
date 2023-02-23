package model

import "time"

type UserSubject struct {
	Id              int       `gorm:"column:user_id" json:"user_id" form:"user_id" db:"user_id"`
	Username        string    `gorm:"column:username" json:"username" form:"username" db:"username"`
	Password        string    `gorm:"column:password" json:"password" form:"password" db:"password"`
	Email           string    `gorm:"column:email" json:"email" form:"email" db:"email"`
	BackgroundImage string    `gorm:"column:background_image" json:"background_image" form:"background_image" db:"background_image"`
	Signature       string    `gorm:"column:signature" json:"signature" form:"signature" db:"signature"`
	Avatar          string    `gorm:"column:avatar" json:"avatar" form:"avatar" db:"avatar"`
	CreateTime      time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time" form:"update_time" db:"update_time"`
}
