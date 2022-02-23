package main

import (
	"fmt"
	"gin/app/blog"
	"gin/app/shop"
	"gin/routers"
)

func main() {
	// 1.创建路由
	//r := gin.Default()
	//// 2.绑定路由规则，执行的函数
	//// gin.Context，封装了request和response
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	fmt.Println(action)
	//	//截取/
	//	action = strings.Trim(action, "/")
	//	c.String(http.StatusOK, name+" is "+action)
	//})
	//
	//r.GET("/demo", func(c *gin.Context) {
	//	//指定默认值
	//	//http://localhost:8080/demo 才会打印出来默认的值
	//	name := c.DefaultQuery("name", "枯藤")
	//	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	//})
	//r := routers.SetupRouter()
	//if err := r.Run(); err != nil {
	//	fmt.Println("startup service failed, err:%v\n", err)
	//}
	//r := gin.Default()
	//routers.LoadBlog(r)
	//if err := r.Run(); err != nil {
	//	fmt.Println("startup service failed, err:%v\n", err)
	//}
	// 加载多个APP的路由配置
	//routers.Include(shop.Routers, blog.Routers)
	routers.Include(blog.Routers, shop.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
}
