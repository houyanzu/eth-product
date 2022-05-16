package database

import (
	"errors"
	"fmt"
	"github.com/houyanzu/eth-product/config"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
)

type MysqlContext struct {
	Db    *gorm.DB
	Error error
}

// SqlDB .
var (
	sqlDB *gorm.DB
)

// InitMysql .
func InitMysql[T any](conf *config.Config[T]) error {
	var err error
	//config := conf.GetConfig()

	param := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.DBName, conf.Mysql.Charset)
	sqlDB, err = gorm.Open(
		mysql.Open(param), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   conf.Mysql.Prefix, // 表名前缀
				SingularTable: true,              // 使用单数表名
			},
			//Logger: logger.Default.LogMode(logger.Info),
		})
	if err != nil {
		return err
	}
	return nil
}

// GetDB .
func GetDB() *gorm.DB {
	if sqlDB == nil {
		panic(errors.New("mysql is not init"))
	}
	return sqlDB
}

func GetContext() *MysqlContext {
	return &MysqlContext{GetDB(), nil}
}

func (m *MysqlContext) Begin() {
	m.Db = m.Db.Begin()
}

func (m *MysqlContext) Commit() {
	m.Db.Commit()
	m.Db = GetDB()
}

func (m *MysqlContext) Rollback() {
	m.Db.Rollback()
	m.Db = GetDB()
}