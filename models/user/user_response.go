package user

import "go-gin-im/models"

type UserResponse struct {
	User models.UserBasic `json:"user"`
}

type LoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	UUID  string `json:"uuid"`
	Token string `json:"token"`
}

// 定义 tag 对应的报错信息
var tagToMsg = map[string]string{
	"eqfield":        "The two passwords must be consistent",
	"aleadyRegister": "The email aleady registered",
	"isRegistered":   "This email has not been registered yet",
}

func GetUserErrMsg(tag string) string {
	msg, ok := tagToMsg[tag]

	if ok {
		return msg
	}

	return ""
}
