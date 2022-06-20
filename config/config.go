package config

// DB .
type Config struct {
	Mysql mysqlConfig `json:"mysql"`
	Redis redisConfig `json:"redis"`
	Eth   ethConfig   `json:"eth"`
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
	LockTransferContract  string `json:"lock_transfer_contract"`
}

var internalConfig *Config

// ParseConfig .
func ParseConfig(conf *Config) {
	internalConfig = conf
}

func GetConfig() *Config {
	return internalConfig
}
