package upload

import "mime/multipart"

type UploadParam struct {
	FileHeader  *multipart.FileHeader
	ContentType string `json:"content_type" form:"content_type" bindind:"required"`
}
