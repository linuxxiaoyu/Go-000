package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linuxxiaoyu/Go-000/Week02/apiv1"
)

func main() {
	r := gin.Default()
	r.GET("/api1/item/:id", apiv1.Item)
	r.Run(":8080")
}
