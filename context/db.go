package context

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var CommonDB *gorm.DB
var DataDB *gorm.DB

func OpenDB(config *Config) (*gorm.DB, error) {
	var err error
	log.Println("正在连接测试数据库... ")
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms", config.GinTest.User, config.GinTest.Password, config.GinTest.Host, config.GinTest.Port, config.GinTest.DbName))

	if err != nil {
		log.Println(err, "连接错误，将在5S内重试... ")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenDB(config)
	}

	if DB.Error != nil {
		log.Println(DB.Error, "连接错误，将在5S内重试... ")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenDB(config)
	}

	log.Println("测试数据库连接成功!")
	DB.LogMode(true)
	DB.SingularTable(true)
	return DB, nil
}

func OpenCommonDB(config *Config) (*gorm.DB, error) {
	var err error
	log.Println("正在连接公共数据库... ")
	CommonDB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms", config.DhbCommon.User, config.DhbCommon.Password, config.DhbCommon.Host, config.DhbCommon.Port, config.DhbCommon.DbName))

	if err != nil {
		log.Println(err, "连接错误，将在5S内重试... ")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenCommonDB(config)
	}

	if CommonDB.Error != nil {
		log.Println(CommonDB.Error, "连接错误，将在5S内重试... ")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenCommonDB(config)
	}

	log.Println("公共数据库连接成功!")
	CommonDB.LogMode(true)
	CommonDB.SingularTable(true)
	return CommonDB, nil
}

func OpenDataDB(config *Config) (*gorm.DB, error) {
	var err error
	log.Println("正在连接业务数据库... ")
	DataDB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms", config.DhbData.User, config.DhbData.Password, config.DhbData.Host, config.DhbData.Port, config.DhbData.DbName))

	if err != nil {
		log.Println(err, "连接错误，将在5S内重试... ")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenDataDB(config)
	}

	if DataDB.Error != nil {
		log.Println(DataDB.Error, "连接错误，将在5S内重试... ")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenDataDB(config)
	}

	log.Println("业务数据库连接成功!")
	DataDB.LogMode(true)
	DataDB.SingularTable(true)
	return DataDB, nil
}
