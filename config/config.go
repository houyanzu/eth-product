package config

import (
	"encoding/json"
	"io/ioutil"
)

// DB .
type Config[T any] struct {
	Mysql  mysqlConfig `json:"mysql"`
	Redis  redisConfig `json:"redis"`
	Common T           `json:"common"`
	Eth    ethConfig   `json:"eth"`
}

type mysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Port     string `json:"port"`
	Prefix   string `json:"prefix"`
	Charset  string `json:"charset"`
}

type redisConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Db       int    `json:"db"`
}

type ethConfig struct {
	Host                  string `json:"host"`
	ChainId               int64  `json:"chain_id"`
	MultiTransferContract string `json:"multi_transfer_contract"`
}

// ParseConfig .
func ParseConfig[T any](fileName string) (res Config[T], err error) {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	err = json.Unmarshal(dat, &res)
	if err != nil {
		return
	}

	return
}

// GetConfig .
//func GetConfig[T any]() *Config[T] {
//	return config
//}

//func CreateConfigFile() {
//	var conf Config
//	js, err := json.Marshal(conf)
//	if err != nil {
//		panic(err)
//	}
//	err = os.WriteFile("config.json", js, 0777)
//	if err != nil {
//		panic(err)
//	}
//}
