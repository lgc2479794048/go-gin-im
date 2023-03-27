package upload

import (
	uploadModel "go-gin-im/models/upload"
	request "go-gin-im/utils/request"
	response "go-gin-im/utils/response"
	"net/http"

	uploadService "go-gin-im/service/upload"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	var params uploadModel.UploadParam
	// 接收文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	params.FileHeader = header
	params.ContentType = header.Header["Content-Type"][0]
	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, request.GetErrMsg(params, err))
		return
	}
	// 调用上传逻辑
	info, err := uploadService.UploadToMinio(c, params)
	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, request.GetErrMsg(params, err))
		return
	}
	response.OK(c, info)
}
