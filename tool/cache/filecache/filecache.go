package filecache

import (
	"time"

	imcache "github.com/houyanzu/cache"
	_ "github.com/houyanzu/cache/redis"
)

var ca imcache.Cache

// InitCache .
func InitCache() error {
	ca = imcache.NewFileCache()
	return nil

}

// Set .
func Set(key string, value interface{}, timeout int64) error {
	if ca == nil {
		_ = InitCache()
	}
	//fmt.Printf("%+v", ca)
	return ca.Put(key, value, time.Duration(timeout)*time.Second)
}

// Get .
func Get(key string) interface{} {
	if ca == nil {
		_ = InitCache()
	}
	return ca.Get(key)
}

func GetString(key string) string {
	if ca == nil {
		_ = InitCache()
	}
	res := ca.Get(key)
	return imcache.GetString(res)
}

func GetInt64(key string) int64 {
	if ca == nil {
		_ = InitCache()
	}
	res := ca.Get(key)
	return imcache.GetInt64(res)
}

// IsExist .
func IsExist(key string) bool {
	if ca == nil {
		_ = InitCache()
	}
	return ca.IsExist(key)
}

// Delete .
func Delete(key string) error {
	if ca == nil {
		_ = InitCache()
	}
	return ca.Delete(key)
}
