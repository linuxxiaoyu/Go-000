package apiv1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/linuxxiaoyu/Go-000/Week02/service"
)

// GET /api1/item/:id
func Item(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
		})
		return
	}

	item, err := service.Item(int(id))
	if err != nil {
		fmt.Printf("Get Item failed: %+v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, item)
}
