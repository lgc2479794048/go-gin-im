/*
 * @Author: lgc2479794048 lgc2479794048@gmail.com
 * @Date: 2023-03-25 19:45:57
 * @LastEditors: lgc2479794048 lgc2479794048@gmail.com
 * @LastEditTime: 2023-03-26 16:56:08
 * @FilePath: \go-gin-im\config\config.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var rootDir = autoMatchDir()

type AppConfig struct {
	App struct {
		Name    string `toml:"name"`
		Version string `toml:"version"`
		Debug   bool   `toml:"debug"`
	} `toml:"app"`
	Server struct {
		Port int    `toml:"port"`
		Mode string `toml:"mode"`
	} `toml:"server"`
}

func NewAppConfig() (*AppConfig, error) {
	config := &AppConfig{}
	data, err := ioutil.ReadFile("app.toml")
	if err != nil {
		return nil, err
	}
	if err := toml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}

type DbConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}
type MinioConfig struct {
	// MinIO服务的端点、访问密钥ID和秘密访问密钥
	EndPoint        string `toml:"endpoint"`
	AccessKeyID     string `toml:"access_key"`
	SecretAccessKey string `toml:"secret_key"`
	UseSSl          bool   `toml:"user_ssl"`
}

func LoadDbConfig(confName string) (*DbConfig, error) {
	config := &DbConfig{}
	filename := filepath.Join(rootDir+"/config/mysql/", fmt.Sprintf("%s.toml", confName))
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := toml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}

func Connect(config *DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewMysqlInstance(serviceName string) (*gorm.DB, error) {
	config, err := LoadDbConfig(serviceName)
	if err != nil {
		log.Fatalf("load config file failed: %v", err)
	}

	db, err := Connect(config)
	if err != nil {
		log.Fatalf("connect database failed: %v", err)
	}
	return db, err
}

func autoMatchDir() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	// 匹配目录,进行替换
	matchStr := "go-gin-im"
	subLen := len(matchStr)
	idx := strings.Index(wd, matchStr)
	if idx >= 0 {
		wd = strings.TrimRight(wd[:idx+subLen], " ")
	}
	fmt.Println(wd)
	return wd
}

func LoadMinioConfig() (*MinioConfig, error) {
	config := &MinioConfig{}
	filename := filepath.Join(rootDir + "/config/minio/minio.toml")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := toml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}

func NewMinioClient() (*minio.Client, error) {
	config, err := LoadMinioConfig()
	if err != nil {
		return nil, err
	}
	minioClient, err := minio.New(config.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSl,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient, nil
}
