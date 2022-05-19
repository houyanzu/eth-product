package chainrecord

import (
	"github.com/houyanzu/eth-product/database"
	"github.com/houyanzu/eth-product/lib/mytime"
	"gorm.io/gorm"
	"strings"
)

type MonitorChainRecord struct {
	ID         uint
	Contract   string
	BlockNum   uint64
	EventId    string
	Hash       string
	CreateTime mytime.DateTime
}

var haveTable = false

func (c *MonitorChainRecord) BeforeCreate(tx *gorm.DB) error {
	c.Contract = strings.ToLower(c.Contract)
	c.CreateTime = mytime.NewFromNow()
	return nil
}

func createTable() error {
	db := database.GetDB()
	return db.Exec("CREATE TABLE `monitor_chain_record` (\n\t`id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,\n\t`contract` char(42) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,\n\t`block_num` int(11) UNSIGNED NOT NULL,\n\t`event_id` char(66) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,\n\t`hash` char(66) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,\n\t`create_time` datetime NOT NULL,\n\tPRIMARY KEY (`id`),\n\tKEY `cb`(`contract`,`block_num`) USING BTREE\n) ENGINE=InnoDB\nDEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci\nAUTO_INCREMENT=1\nROW_FORMAT=DYNAMIC\nAVG_ROW_LENGTH=0;").Error
}

type Model struct {
	*database.MysqlContext
	Data  MonitorChainRecord
	List  []MonitorChainRecord
	Total int64
}

func New(ctx *database.MysqlContext) *Model {
	if ctx == nil {
		ctx = database.GetContext()
	}
	list := make([]MonitorChainRecord, 0)
	data := MonitorChainRecord{}
	if !haveTable {
		hasTable := ctx.Db.Migrator().HasTable(&data)
		if !hasTable {
			err := createTable()
			if err != nil {
				panic(err)
			}
		}
		haveTable = true
	}

	return &Model{ctx, data, list, 0}
}

func (m *Model) InitByUserData(data MonitorChainRecord) *Model {
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

func GetLastBlockNum(contract string) uint64 {
	db := database.GetDB()
	var c MonitorChainRecord
	db.Where("contract = ?", contract).Order("block_num desc").Take(&c)
	return c.BlockNum
}
