package context

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func OpenDB(config *Config) (*gorm.DB, error) {
	var err error
	log.Println("正在连接数据库... ")
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

	log.Println("数据库连接成功!")
	return DB, nil
}
