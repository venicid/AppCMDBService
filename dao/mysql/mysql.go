package mysql

import (
	"AppCMDBService/logger"
	"AppCMDBService/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 驱动mysql
)

var (
	isInit bool
	GORM   *gorm.DB
	err    error
)

// db的初始化函数，与数据库建立连接
func Init() {
	// 判断是否已经初始化了
	if isInit {
		return
	}

	// 组装连接配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.Conf.MySQLConfig.User,
		setting.Conf.MySQLConfig.Password,
		setting.Conf.MySQLConfig.Host,
		setting.Conf.MySQLConfig.Port,
		setting.Conf.MySQLConfig.DB,
	)
	fmt.Println(dsn)
	GORM, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("db conn  failed" + err.Error())
	}

	// 打印sql语句
	GORM.LogMode(setting.Conf.MySQLConfig.LogModel)

	// 开启连接池
	GORM.DB().SetMaxIdleConns(setting.Conf.MySQLConfig.MaxIdleConns)
	GORM.DB().SetMaxIdleConns(setting.Conf.MySQLConfig.MaxOpenConns)
	//GORM.DB().SetConnMaxLifetime(time.Duration(config.MaxLifeTime))

	isInit = true
	logger.Logger.Info("db conn success")
}

// 关闭数据库连接
func Close() error {
	return GORM.Close()
}
