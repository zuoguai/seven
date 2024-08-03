package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

const (
	DEFAULT_CONFIG_PATH = ".././app.yaml"
)

type Configs struct {
	App          AppConfig     `json:"app" yaml:"app"`
	Mysql        MysqlConfig   `json:"mysql" yaml:"mysql"`
	PrivateToken PrivateConfig `json:"private_token" yaml:"private_token"`
}

var (
	configs *Configs
	once    sync.Once
)

type MysqlConfig struct {
	Database string `json:"database" yaml:"database"`
	User     string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
}
type AppConfig struct {
	AppName    string `json:"app_name" yaml:"app_name"`
	AppVersion string `json:"app_version" yaml:"app_version"`
	ServerPort int    `json:"server_port" yaml:"server_port"`
}
type PrivateConfig struct {
	ApiToken string `json:"api_token" yaml:"api_token"`
}

func GetConfigs(path string) *Configs {

	once.Do(func() {
		if path == "" {
			path = DEFAULT_CONFIG_PATH
		}
		InitConfig(path)
	})
	return configs
}

func InitConfig(path string) {
	fd, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()

	content, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println(err)

	}

	if strings.HasSuffix(path, ".json") {
		if err = jsoniter.Unmarshal(content, &configs); err != nil {
			fmt.Println(err)

		}
	} else if strings.HasSuffix(path, ".yaml") {
		if err = yaml.Unmarshal(content, &configs); err != nil {
			fmt.Println(err)

		}
	}
}
