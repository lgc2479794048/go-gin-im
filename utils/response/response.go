/*
 * @Author: lgc2479794048 lgc2479794048@gmail.com
 * @Date: 2023-03-26 00:42:36
 * @LastEditors: lgc2479794048 lgc2479794048@gmail.com
 * @LastEditTime: 2023-03-26 00:43:06
 * @FilePath: \go-gin-im\utils\response\response.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

var defaultMsg = map[int]string{
	http.StatusOK: "Success",

	http.StatusUnauthorized:        "User not logged in",
	http.StatusNotFound:            "Not Found",
	http.StatusUnprocessableEntity: "Data was wrong",
}

func Result(status int, data interface{}, c *gin.Context, msg []string) {
	returnMsg := defaultMsg[status]
	if len(msg) == 1 {
		returnMsg = msg[0]
	} else if returnMsg == "" {
		returnMsg = "something was wrong"
	}

	c.JSON(status, Response{
		data,
		returnMsg,
	})
}

func OK(c *gin.Context, data interface{}, msg ...string) {
	Result(http.StatusOK, data, c, msg)
}

func Failed(c *gin.Context, status int, msg ...string) {
	Result(status, map[string]interface{}{}, c, msg)
}
