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
	Mysql MysqlConfig `json:"mysql" yaml:"mysql"`
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

func GetConfigs() *Configs {
	once.Do(func() {
		InitConfig(DEFAULT_CONFIG_PATH)
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
