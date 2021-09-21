package models

import "time"

type User struct {
	ID             int       `gorm:"column:user_id;AUTO_INCREMENT" form:"id" json:"id"`
	Email          string    `gorm:"column:email" form:"email" json:"email"`
	Name           string    `gorm:"column:name" form:"name" json:"name"`
	Gender         string    `gorm:"column:gender" form:"gender" json:"gender"`
	Password       string    `gorm:"column:password" form:"password"`
	CPassword      string    `gorm:"-" form:"cpassword"`
	TokenForgot    string    `gorm:"column:token_forgot" form:"token_forgot"`
	TokenForgotExp time.Time `gorm:"column:token_forgot_expired_time"`
	Isdelete       bool      `gorm:"column:is_delete" json:"is_delete"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	Photo          string    `gorm:"column:photo" json:"photo"`
}
