package shop

import (
	_ "gin/config"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/getNum", postHandler)
	e.GET("/getValue", getValue)
}
