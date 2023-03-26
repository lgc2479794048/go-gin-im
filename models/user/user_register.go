/*
 * @Author: lgc2479794048 lgc2479794048@gmail.com
 * @Date: 2023-03-25 23:57:42
 * @LastEditors: lgc2479794048 lgc2479794048@gmail.com
 * @LastEditTime: 2023-03-25 23:59:07
 * @FilePath: \go-gin-im\models\user\user_register.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package user

type UserRegisterParam struct {
	Name          string `json:"name" example:"name" binding:"required"`
	Password      string `json:"password" example:"password" binding:"required"`
	ConfirmedPass string `json:"confirmed password" example:"password" binding:"required,eqfield=Password"`
	Email         string `json:"email"  binding:"required"`
}

type UserLoginParam struct {
	Password string `json:"password" example:"password" binding:"required"`
	Email    string `json:"email"  binding:"required"`
}
