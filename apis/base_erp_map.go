package apis

import (
	model "gin_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Categorys 列表数据
func GetMaps(c *gin.Context) {
	var maps model.BaseErpMap

	result, err := maps.BaseErpMaps()

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
