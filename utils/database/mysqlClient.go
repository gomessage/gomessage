package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// mysqlClient 默认数据库
func mysqlClient() *gorm.DB {
	//return openMysqlConnector("deploy_engine", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx:3306", "xxxxxxxxxxxxx", "xxxxxxxxxxxxx", "charset=utf8mb4&parseTime=True&loc=Local")
	return openMysqlConnector(
		viper.GetString("mysql.host")+":"+viper.GetString("mysql.port"),
		viper.GetString("mysql.name"),
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.options"),
	)
}

// Mysql的数据库连接
func openMysqlConnector(databaseURL, databaseName, username, password, options string) *gorm.DB {

	//dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, databaseURL, databaseName)
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", username, password, databaseURL, databaseName, options)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	//创建一个连接池
	sqlDB, err := database.DB()
	if err != nil {
		fmt.Println(err)
	}

	//SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	//SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(10000)
	//SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	//返回一个数据库实例对象
	return database
}
