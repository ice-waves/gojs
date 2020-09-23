package db

import (
	"github.com/ice-waves/gojs/configs/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once

// 单例模式
func init() {
	once.Do(func() {
		InitDb()
	})
}

func InitDb() *gorm.DB {
	databaseConf := new(conf.Database).GetConf()
	db, err := gorm.Open(mysql.Open(databaseConf.Mysql.DSN), &gorm.Config{})
	if err != nil {
		panic("mysql connect failed!")
	}

	return db
}
