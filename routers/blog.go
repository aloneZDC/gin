package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadBlog(e *gin.Engine) {
	e.GET("/post", postHandler)
	e.Run(":8000")
}

func postHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.xxx.com",
	})
}
