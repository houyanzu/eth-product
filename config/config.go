package config

import (
	"encoding/json"
	"io/ioutil"
)

// DB .
type Config struct {
	Mysql mysqlConfig `json:"mysql"`
	Redis redisConfig `json:"redis"`
	Eth   ethConfig   `json:"eth"`
	Extra extra       `json:"extra"`
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
	Prefix   string `json:"prefix"`
}

type ethConfig struct {
	Host                  string `json:"host"`
	ApiHost               string `json:"api_host"`
	ApiKey                string `json:"api_key"`
	ChainId               int64  `json:"chain_id"`
	MultiTransferContract string `json:"multi_transfer_contract"`
	LockTransferContract  string `json:"lock_transfer_contract"`
}

type extra struct {
	DingTalkURL string `json:"ding_talk_url"`
}

var internalConfig *Config

// ParseConfig .
func ParseConfig(conf *Config) {
	internalConfig = conf
}

func ParseConfigByFile(configName string) error {
	dat, err := ioutil.ReadFile(configName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dat, &internalConfig)
	if err != nil {
		return err
	}
	return nil
}

func GetConfig() *Config {
	return internalConfig
}
