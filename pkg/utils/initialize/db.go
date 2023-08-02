package initialize

import (
	"gomessage/pkg/authorization"
	"gomessage/pkg/models"
	"gomessage/pkg/utils/database"
	"gomessage/pkg/utils/log/loggers"
	"gorm.io/gorm"
)

// InitDB 初始化函数
func InitDB(databaseType string, isMigrate bool) {
	//针对指定环境进行数据库初始化
	database.DB.Init(databaseType)
	loggers.DefaultLogger.Info("数据库初始化完成...")

	//判断是否要迁移一次数据库表结构
	isAutoMigrateDB(isMigrate)
}

// 判断是否需要自动迁移
func isAutoMigrateDB(migrate bool) {
	if migrate {
		//迁移数据库表结构（把需要迁移的表结构体，追加到下文括号中，有点类似于Django注册app的感觉）
		var DbList []any
		DbList = append(DbList,
			&models.Namespace{},
			&models.Template{},
			&models.Variables{},
			&models.Client{},
			&models.Dingtalk{},
			&models.Feishu{},
			&models.WechatApplication{},
			&models.WechatRobot{},
			&authorization.Users{},
			&authorization.Sessions{},
		)

		//数据库自动迁移
		//参数1：如果同时连接了多个数据库，需要明确指定往哪个客户端迁移，这里默认往database.DB.Client客户端对应的数据库中进行迁移
		//参数2：这是一个list类型，需要指定要进行迁移的有哪些表
		databaseAutoMigrate(database.DB.Default, DbList)
		loggers.DefaultLogger.Info("检测到 '--migrate==True' 开始对数据库进行幂等性检测...")
	}

	//创建默认的Namespace（只有真正迁移数据库时，才会创建default命名空间）
	models.InitNamespace()
	//创建默认的admin账户（只有真正迁移数据库时，才会初始化创建admin账户，若admin已存在则跳过初始化）
	authorization.InitAdmin()
}

// 具体的迁移方法
func databaseAutoMigrate(databaseClient *gorm.DB, databaseList []any) {
	for _, table := range databaseList {
		//如果不做判断，那么gorm每一次都是增量式迁移，字段只增不见，会保留无用的表字段这样比较安全
		if err := databaseClient.AutoMigrate(&table); err != nil {
			loggers.DefaultLogger.Warningln("数据库迁移失败...：%s", table)
			return
		}
	}
}
