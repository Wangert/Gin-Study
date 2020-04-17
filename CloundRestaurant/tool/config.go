package tool

import (
	"os"
	"bufio"
	"encoding/json"
)

//配置结构体
type Config struct {
	AppName  string `json:"app_name"`
	AppMode  string `json:"app_mode"`
	AppHost  string `json:"app_host"`
	AppPort  string `json:"app_port"`
	Database DatabaseConfig `json:"database"`
	Redis    RedisConfig `json:"redis"`
}

//数据库配置结构体
type DatabaseConfig struct {
	Driver	 string	`json:"driver"`
	User	 string	`json:"user"`
	Password string `json:"password"`
	Host 	 string	`json:"host"`
	Port 	 string	`json:"port"`
	DBName 	 string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool	`json:"show_sql"`
}

//redis参数配置
type RedisConfig struct {
	Addr	 string	`json:"addr"`
	Port	 string	`json:"port"`
	Password string	`json:"password"`
	Db 		 int 	`json:"db"`
}

//配置对象
var cfg Config

func ParseConfig(path string) (*Config, error) {
	//打开配置文件
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer configFile.Close()
	//创建文件对应reader
	reader := bufio.NewReader(configFile)
	//创建json解码器
	jsonDecoder := json.NewDecoder(reader)
	//进行解码获取配置文件内容到配置对象
	if err = jsonDecoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func GetConfig() *Config {
	return &cfg
}