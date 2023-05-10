package common

import (
	"bufio"
	"fmt"
	"os"
)

type Config struct {
	Port     int32 `json:"port"`
	Database struct {
		Port     int32  `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"database"`
}

var config Config

func ReadConfig() {
	file, err := os.Open("./config/config.json")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var str string
	for scanner.Scan() {
		str += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件出错：", err)
		return
	}
	JSONParse(str, &config)
}

func GetConfig() *Config {
	return &config
}
