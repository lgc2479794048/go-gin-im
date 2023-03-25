package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	UserName string
	Sex      uint8
	Profile  string
	Age      int
	DbTime   int64
}
