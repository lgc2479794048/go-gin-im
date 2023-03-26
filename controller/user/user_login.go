package user

import (
	"fmt"
	"go-gin-im/config"
	"go-gin-im/constant"
	"go-gin-im/models"
	userModel "go-gin-im/models/user"
	request "go-gin-im/utils/request"
	response "go-gin-im/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var params userModel.UserLoginParam
	var user models.UserBasic

	// 验证参数
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, request.GetErrMsg(params, err))
		return
	}
	user = models.UserBasic{}
	fmt.Println(user, err)
	// 进行数据库查询
	db, err := config.NewMysqlInstance(constant.DB_GO_GIN_IM_SERVICE_NAME)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Where("email=?", params.Email).Find(&user).Error
	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, "登录失败")
		return
	}

	response.OK(c, userModel.UserResponse{User: user})
}
