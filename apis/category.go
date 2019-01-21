package apis

import (
	model "gin_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetCategories 列表数据
func GetCategories(c *gin.Context) {
	var category model.BaseCategory

	result, err := category.GetCategories()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    101,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  100,
		"count": len(result),
		"data":  result,
	})
}

//GetCategories 列表数据
func GetSites(c *gin.Context) {
	var category model.BaseCategory

	result, err := category.GetSite()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    101,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  100,
		"count": len(result),
		"data":  result,
	})
}

//GetCategories 列表数据
func GetSites2(c *gin.Context) {
	var category model.BaseCategory

	result, err := category.GetSite2()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    101,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  100,
		"count": len(result),
		"data":  result,
	})
}
