package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	UserName string
	Sex      uint8
	Profile  string
	DbTime   int64 `json:"db_time" gorm:"required"`
	Age      int
	Password string `json:"-" gorm:"required"`
	Email    string `json:"email" gorm:"required,unique"`
	UUID     string `json:"uuid" gorm:"required,unique"`
	ClientIP string `json:"clientIP" gorm:"required"`
}
