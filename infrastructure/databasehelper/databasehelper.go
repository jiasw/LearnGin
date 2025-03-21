package databasehelper

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DatabaseHelper struct {
	DB *gorm.DB
}

func NewDatabaseHelper(dsn string) (*DatabaseHelper, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 生产环境建议改为 logger.Error
	})
	if err != nil {
		return nil, err
	}
	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(50) // 最大连接数
	sqlDB.SetMaxIdleConns(20) // 空闲连接数
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	return &DatabaseHelper{DB: db}, nil
}

func (db *DatabaseHelper) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
