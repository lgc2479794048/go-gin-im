package user

import (
	"fmt"
	"go-gin-im/config"
	"go-gin-im/constant"
	"go-gin-im/models"
	userModel "go-gin-im/models/user"
	"go-gin-im/utils"
	request "go-gin-im/utils/request"
	response "go-gin-im/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Register(c *gin.Context) {
	var params userModel.UserRegisterParam
	var user models.UserBasic

	// 验证参数
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, request.GetErrMsg(params, err))
		return
	}

	// 生成uuid
	u1, err := uuid.NewRandom()
	for err != nil {
		u1, err = uuid.NewRandom()
	}

	user = models.UserBasic{
		UserName: params.Name,
		Password: utils.BcryptHash(params.Password),
		Email:    params.Email,
		ClientIP: c.ClientIP(),
		UUID:     u1.String(),
	}

	fmt.Println(user, err)
	// 添加到数据库
	db, err := config.NewMysqlInstance(constant.DB_GO_GIN_IM_SERVICE_NAME)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 先判断是否已经注册过
	account := models.UserBasic{}
	err = db.Where("email=?", params.Email).Find(&account).Error
	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, "注册失败")
		return
	}
	if account.ID != 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "该邮箱已经注册")
		return
	}
	err = db.Create(&user).Error
	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, "注册失败")
		return
	}

	response.OK(c, userModel.UserResponse{User: user})
}
