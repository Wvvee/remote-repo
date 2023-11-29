package dao

import (
	"gin-ranking/config"
	"gin-ranking/pkg/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// "github.com/sirupsen/logrus"
)

var (
	Db  *gorm.DB
	err error
)

func init() {

	Db, err = gorm.Open("mysql", config.Mysqldb)
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
		return
	}
	if Db.Error != nil {
		logger.Error(map[string]interface{}{"database error": Db.Error})
		return
	}

	//设置连接池
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour)

	// //测试数据库是否可以连接
	// if err := Db.DB().Ping(); err != nil {
	// 	logger.Error(map[string]interface{}{"database ping error": err.Error()})
	// } else {
	// 	logger.Info(logrus.Fields{"message": "Database connection successful."})
	// }
}
