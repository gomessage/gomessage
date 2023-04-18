package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path"
)

func sqlite3Client() *gorm.DB {
	return openSqlite3Connector()
}

// sqlite3的数据库连接
func openSqlite3Connector() *gorm.DB {

	dataPath := viper.GetString("sqlite3.path")
	//确保存放日志文件的目录始终存在
	dataPathDir := path.Dir(dataPath)                             //返回路径中除去最后一个元素的剩余部分，也就是路径最后一个元素所在的目录
	if err := os.MkdirAll(dataPathDir, os.ModePerm); err != nil { //创建目录类似于（mkdir -p /aaa/bbb的效果）
		fmt.Println("创建目录失败：", err)
		os.Exit(3)
	}

	database, err := gorm.Open(sqlite.Open(dataPath), &gorm.Config{
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
