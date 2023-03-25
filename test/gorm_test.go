/*
 * @Author: lgc2479794048 lgc2479794048@gmail.com
 * @Date: 2023-03-25 14:47:34
 * @LastEditors: lgc2479794048 lgc2479794048@gmail.com
 * @LastEditTime: 2023-03-25 14:53:40
 * @FilePath: \go-gin-im\test\gorm_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package test

import (
	"fmt"
	"gp-gin-im/models"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGorm(t *testing.T) {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模型
	db.AutoMigrate(&models.UserBasic{})

	// 插入数据
	user := models.UserBasic{
		UserName: "lingengcheng",
		Sex:      1,
	}
	result := db.Create(&user)
	if result.Error != nil {
		panic("failed to insert data")
	}

	// 查询数据
	var users []models.UserBasic
	db.Find(&users)
	fmt.Println(users)

	// 更新数据
	db.Model(&user).Update("Sex", 2)
	db.Find(&users)
	fmt.Println(users)

	// 删除数据
	db.Delete(&user)
	db.Find(&users)
	fmt.Println(users)
}
