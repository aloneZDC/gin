package routers

import (
	"github.com/gin-gonic/gin"
)

//func helloHandler(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"message": "Hello www.topgoer.com",
//	})
//}
//
//// SetupRouter 配置路由信息
//func SetupRouter() *gin.Engine {
//	r := gin.Default()
//	r.GET("/", helloHandler)
//	r.Run(":8000")
//	return r
//}

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	r.Run(":8000")
	return r
}
