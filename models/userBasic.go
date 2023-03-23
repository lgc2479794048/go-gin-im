package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	UserName string
	Sex      uint8
}
