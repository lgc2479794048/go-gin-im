package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func LoadDbConfig(confName string) (*DbConfig, error) {
	config := &DbConfig{}
	data, err := ioutil.ReadFile(fmt.Sprintf("../config/mysql/%s.toml", confName))
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
