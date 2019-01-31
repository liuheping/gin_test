package router

import (
	. "gin_test/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/getCategories", GetCategories)

	router.POST("/maps", GetMaps)

	router.POST("/getSites", GetSites)

	router.POST("/getSites2", GetSites2)

	return router
}
