package database

import "gorm.io/gorm"

var DB *databaseClientList

// 数据库列表，可以同时连接多个数据库
type databaseClientList struct {
	Default *gorm.DB //通用客户端，给它赋值不同的数据库对象（比如MySQL或者SQLite），可以最小限度的修改全局其它地方的代码。
	Mysql   *gorm.DB //Mysql客户端
	Sqlite3 *gorm.DB //Sqlite3客户端
}

func (d *databaseClientList) Init(databaseType string) {
	switch databaseType {
	case "mysql":
		DB = &databaseClientList{
			Default: mysqlClient(), //通用客户端，给它赋值不同的数据库对象（比如MySQL或者SQLite），可以不用修改全局其它地方的代码
			Mysql:   mysqlClient(), //Mysql专用客户端
		}

	case "sqlite", "sqlite3":
		DB = &databaseClientList{
			Default: sqlite3Client(), //通用客户端，给它赋值不同的数据库对象（比如MySQL或者SQLite），可以不用修改全局其它地方的代码
			Sqlite3: sqlite3Client(), //Sqlite3专用客户端
		}

	default:
		DB = &databaseClientList{
			Default: sqlite3Client(), //通用客户端，给它赋值不同的数据库对象（比如MySQL或者SQLite），可以不用修改全局其它地方的代码
			Mysql:   mysqlClient(),   //Mysql专用客户端
			Sqlite3: sqlite3Client(), //Sqlite3专用客户端
		}
	}
}
