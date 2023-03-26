/*
 * @Author: lgc2479794048 lgc2479794048@gmail.com
 * @Date: 2023-03-26 21:20:09
 * @LastEditors: lgc2479794048 lgc2479794048@gmail.com
 * @LastEditTime: 2023-03-26 22:07:08
 * @FilePath: \go-gin-im\test\upload_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func TestUpload(t *testing.T) {
	// 初始化Minio客户端
	endpoint := "localhost:9071"
	accessKey := "lingengcheng"
	secretKey := "12345678"
	useSSL := false
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
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

	// 上传本地文件
	localFilePath := "C:/Users/24797/Pictures/4781442-d6a8c2e5714b4c44.png"
	objectName := "image.jpg"
	contentType := "image/jpeg"
	f, err := os.Open(localFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, f, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("File '%s' uploaded successfully.\n", objectName)
}
