package context

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var CommonDB *gorm.DB
var DataDB *gorm.DB

func OpenCommonDB(config *Config) (*gorm.DB, error) {
	var err error
	log.Println("正在连接公共数据库... ")
	CommonDB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms", config.DhbCommon.User, config.DhbCommon.Password, config.DhbCommon.Host, config.DhbCommon.Port, config.DhbCommon.DbName))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("公共数据库连接成功!")
	CommonDB.LogMode(true)       //启用Logger，显示详细日志
	CommonDB.SingularTable(true) //全局设置表名不可以为复数形式
	return CommonDB, nil
}

func OpenDataDB(config *Config) (*gorm.DB, error) {
	var err error
	log.Println("正在连接业务数据库... ")
	DataDB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms", config.DhbData.User, config.DhbData.Password, config.DhbData.Host, config.DhbData.Port, config.DhbData.DbName))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("业务数据库连接成功!")
	DataDB.LogMode(true)       //启用Logger，显示详细日志
	DataDB.SingularTable(true) //全局设置表名不可以为复数形式
	return DataDB, nil
}
