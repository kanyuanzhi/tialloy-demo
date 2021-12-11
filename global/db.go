package global

import (
	"fmt"
	"github.com/kanyuanzhi/tialloy/tilog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		CustomObject.MysqlUsername,
		CustomObject.MysqlPassword,
		CustomObject.MysqlHost,
		CustomObject.MysqlPort,
		CustomObject.MysqlDbname)
	globalDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		tilog.Log.Error(err)
		return nil
	}
	sqlDB, err := globalDB.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	tilog.Log.Infof("connect to mysql %s:%d, dbname=%s successfully",
		CustomObject.MysqlHost,
		CustomObject.MysqlPort,
		CustomObject.MysqlDbname)

	return globalDB
}
