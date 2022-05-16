package transferrecords

import (
	"github.com/houyanzu/eth-product/database"
	"github.com/houyanzu/eth-product/lib/mytime"
	"gorm.io/gorm"
)

type TransferRecords struct {
	ID         uint
	Hash       string
	Status     int8
	Nonce      uint64
	CreateTime mytime.DateTime
}

func (c *TransferRecords) BeforeCreate(tx *gorm.DB) error {
	c.Status = 1
	c.CreateTime = mytime.NewFromNow()
	return nil
}

func createTable() error {
	db := database.GetDB()
	return db.Exec("CREATE TABLE `transfer_records` (\n\t`id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,\n\t`hash` char(66) NOT NULL,\n\t`nonce` int(11) UNSIGNED NOT NULL,\n\t`status` tinyint(1) NOT NULL,\n\t`create_time` datetime NOT NULL,\n\tPRIMARY KEY (`id`)\n) ENGINE=InnoDB\nDEFAULT CHARACTER SET=utf8;").Error
}

type Model struct {
	*database.MysqlContext
	Data  TransferRecords
	List  []TransferRecords
	Total int64
}

func New(ctx *database.MysqlContext) *Model {
	if ctx == nil {
		ctx = database.GetContext()
	}
	list := make([]TransferRecords, 0)
	data := TransferRecords{}
	hasTable := ctx.Db.Migrator().HasTable(&data)
	if !hasTable {
		err := createTable()
		if err != nil {
			panic(err)
		}
	}
	return &Model{ctx, data, list, 0}
}

func (m *Model) InitByUserData(data TransferRecords) *Model {
	m.Data = data
	return m
}

func (m *Model) Foreach(f func(index int, m *Model)) {
	for k, v := range m.List {
		mm := New(nil).InitByUserData(v)
		f(k, mm)
	}
}

func (m *Model) Add() {
	m.Error = m.Db.Create(&m.Data).Error
}

func (m *Model) InitPending() *Model {
	m.Error = m.Db.Where("status = 1").Take(&m.Data).Error
	return m
}

func (m *Model) SetSuccess() bool {
	m.Error = m.Db.Model(&m.Data).Updates(map[string]any{
		"status": 2,
	}).Error
	return true
}

func (m *Model) SetFail() bool {
	m.Error = m.Db.Model(&m.Data).Updates(map[string]any{
		"status": -1,
	}).Error
	return true
}
