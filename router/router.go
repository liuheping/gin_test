package router

import (
	. "gin_test/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getCategories", GetCategories)

	router.GET("/maps", GetMaps)

	router.GET("/getSites", GetSites)

	router.GET("/getSites2", GetSites2)

	return router
}
