package model

import "time"

type MessageSubject struct {
	Id         int       `gorm:"column:message_id" json:"message_id" form:"message_id" db:"message_id"`
	FromUserId int       `gorm:"column:from_user_id" json:"from_user_id" form:"from_user_id" db:"from_user_id"`
	Message    string    `gorm:"column:message" json:"message" form:"message" db:"message"`
	ToUserId   int       `gorm:"column:to_user_id" json:"to_user_id" form:"to_user_id" db:"to_user_id"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time" form:"create_time" db:"create_time"`
}
