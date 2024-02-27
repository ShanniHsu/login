package model

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"column:id" json:"id"`
	LastName  string    `gorm:"column:last_name" json:"last_name"`
	FirstName string    `gorm:"column:first_name" json:"first_name"`
	NickName  string    `gorm:"column:nick_name" json:"nick_name"`
	Account   string    `gorm:"column:account" json:"account"`
	Password  string    `gorm:"column:password" json:"password"`
	Email     string    `gorm:"column:email" json:"email"`
	Gender    string    `gorm:"column:gender" json:"gender"`
	Amount    int64     `gorm:"column:amount" json:"amount"`
	Token     string    `gorm:"column:token" json:"token"`
	TempToken string    `gorm:"column:temp_token" json:"temp_token"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
