package dao

import (
	"errors"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/penguinn/go-sdk/log"
	"github.com/penguinn/goframe/config"
	"github.com/penguinn/goframe/constant"
)

var dbMap = sync.Map{}

func Init() error {
	for name, dsn := range config.Config.DSN {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil || db == nil {
			log.Error(err)
			return errors.New("无法从平台服务获取数据库连接" + err.Error())
		}
		conn, err := db.DB()
		if err != nil {
			log.Error(err)
			return errors.New("无法从平台服务获取数据库连接" + err.Error())
		}
		conn.SetMaxIdleConns(10)
		conn.SetMaxOpenConns(40)

		SetDB(name, db)
		log.Debugf("初始化数据库[%s]成功", name)
	}
	return nil
}

func SetDB(dbName string, db *gorm.DB) {
	dbMap.Store(dbName, db)
}

func GetDefault() *gorm.DB {
	return GetDB(constant.DBDefault)
}

func GetDB(dbName string) *gorm.DB {
	db, ok := dbMap.Load(dbName)
	if !ok {
		db, _ = dbMap.Load(constant.DBDefault)
	}

	return db.(*gorm.DB)
}
