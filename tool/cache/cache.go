package cache

import (
	"errors"
	"fmt"
	"github.com/houyanzu/eth-product/config"
	"strconv"
	"time"

	imcache "github.com/houyanzu/cache"
	_ "github.com/houyanzu/cache/redis"
)

var ca imcache.Cache

// InitCache .
func InitCache[T any](conf config.Config[T]) error {
	redisConfig := `{
		"conn": "redis://` + conf.Redis.Password + `@` + conf.Redis.Host + `:` + conf.Redis.Port + `",
		"dbNum": "` + fmt.Sprintf("%v", conf.Redis.Db) + `",
		"key": ""
	}`
	//fmt.Println(redisConfig)
	//fmt.Println(redisConfig)
	var err error
	ca, err = imcache.NewCache("redis", redisConfig)
	if err != nil {
		return err
	}
	return nil
}

// Set .
func Set(key string, value interface{}, timeout int64) error {
	if ca == nil {
		panic("cache 未初始化")
	}
	//fmt.Printf("%+v", ca)
	return ca.Put(key, value, time.Duration(timeout)*time.Second)
}

func SetNX(key string, timeout int64) (bool, error) {
	if ca == nil {
		panic("cache 未初始化")
	}
	return ca.SetNX(key, time.Duration(timeout)*time.Second)
}

// Get .
func Get(key string) interface{} {
	if ca == nil {
		panic("cache 未初始化")
	}
	return ca.Get(key)
}

func GetInt64(key string) (int64, error) {
	if ca == nil {
		panic("cache 未初始化")
	}
	resByte, ok := ca.Get(key).([]byte)
	if !ok {
		return 0, errors.New("wrong")
	}
	res, err := strconv.ParseInt(string(resByte), 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func GetString(key string) (string, error) {
	if ca == nil {
		panic("cache 未初始化")
	}
	resByte, ok := ca.Get(key).([]byte)
	if !ok {
		return "", errors.New("wrong")
	}
	return string(resByte), nil
}

func GetUint64(key string) (uint64, error) {
	if ca == nil {
		panic("cache 未初始化")
	}
	resByte, ok := ca.Get(key).([]byte)
	if !ok {
		return 0, errors.New("wrong")
	}
	res, err := strconv.ParseUint(string(resByte), 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func IncrValue(key string, value interface{}) error {
	if ca == nil {
		panic("cache 未初始化")
	}
	return ca.IncrValue(key, value)
}

func DecrValue(key string, value interface{}) error {
	if ca == nil {
		panic("cache 未初始化")
	}
	return ca.DecrValue(key, value)
}

// IsExist .
func IsExist(key string) bool {
	if ca == nil {
		panic("cache 未初始化")
	}
	return ca.IsExist(key)
}

// Delete .
func Delete(key string) error {
	if ca == nil {
		panic("cache 未初始化")
	}
	return ca.Delete(key)
}

func lock(key string) {

}
