package databasehelper

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	instance *DatabaseHelper
	once     sync.Once
)

type DatabaseHelper struct {
	DB *gorm.DB
}

// 初始化方法（只需调用一次）
func Initialize(dsn string) error {
	var initErr error
	once.Do(func() {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			initErr = err
			return
		}
		instance = &DatabaseHelper{DB: db}
	})
	return initErr
}

// 获取单例实例
func GetInstance() *DatabaseHelper {
	return instance
}

func (db *DatabaseHelper) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
