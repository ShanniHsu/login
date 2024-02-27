package model

import (
	"time"
)

type UserLoginLog struct {
	ID        int64     `gorm:"column:id" json:"id"`
	UserID    int64     `gorm:"column:user_id" json:"user_id"`
	Account   string    `gorm:"column:account" json:"account"`
	Result    string    `gorm:"column:result" json:"result"`
	Remark    string    `gorm:"column:remark" json:"remark"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
