package main

import (
	gcontext "gin_test/context"
	"gin_test/router"
	"log"
)

func main() {
	config := gcontext.LoadConfig(".")
	DB, err := gcontext.OpenDB(config)
	if err != nil {
		log.Fatalf("连接数据库失败: %s \n", err)
	}
	CommonDB, err := gcontext.OpenCommonDB(config)
	if err != nil {
		log.Fatalf("连接数据库失败: %s \n", err)
	}
	DataDB, err := gcontext.OpenDataDB(config)
	if err != nil {
		log.Fatalf("连接数据库失败: %s \n", err)
	}

	defer DB.Close()
	defer CommonDB.Close()
	defer DataDB.Close()
	router := router.InitRouter()
	router.Run(":8000")
}
