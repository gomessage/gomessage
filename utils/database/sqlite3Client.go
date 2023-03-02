package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func sqlite3Client() *gorm.DB {
	return openSqlite3Connector()
}

// Mysql的数据库连接
func openSqlite3Connector() *gorm.DB {
	database, err := gorm.Open(sqlite.Open(viper.GetString("sqlite3.path")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), //设置日志默认模式为Silent
	})
	if err != nil {
		fmt.Println(err)
	}

	//创建一个连接池
	sqlDB, err := database.DB()
	if err != nil {
		fmt.Println(err)
	}

	//SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(viper.GetInt("sqlite3.MaxIdleConns"))
	//SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(viper.GetInt("sqlite3.MaxOpenConns"))

	return database
}
