/*
 * @Author: lgc2479794048 lgc2479794048@gmail.com
 * @Date: 2023-03-26 21:09:44
 * @LastEditors: lgc2479794048 lgc2479794048@gmail.com
 * @LastEditTime: 2023-03-26 21:15:34
 * @FilePath: \go-gin-im\service\upload\upload.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package upload

import (
	"context"
	"go-gin-im/config"
	"go-gin-im/models/upload"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func UploadToMinio(c *gin.Context, param upload.UploadParam) (minio.UploadInfo, error) {
	minioClient, err := config.NewMinioClient()
	if err != nil {
		log.Fatalln(err)
	}

	// 创建一个新的存储桶
	bucketName := "go-gin-im"
	location := "us-east-1"
	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// 如果存储桶已经存在，则忽略错误
		if exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName); errBucketExists == nil && exists {
			log.Printf("Bucket '%s' already exists.\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Bucket '%s' created successfully.\n", bucketName)
	}

	contentType := param.ContentType
	// 文件名用uuid作为前缀
	uuid, _ := uuid.NewRandom()
	// 获取文件后缀名
	ext := filepath.Ext(param.FileHeader.Filename)
	objectName := gconv.String(uuid) + ext
	f, err := param.FileHeader.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	info, err := minioClient.PutObject(context.Background(), bucketName, objectName, f, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("File '%s' uploaded successfully.\n", objectName)
	return info, nil
}
