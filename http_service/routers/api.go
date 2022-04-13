package routers

import (
	"github.com/gin-gonic/gin"
	"xt-service/http_service/controllers"
)

func registerApi(g *gin.Engine)  {

	v1 := g.Group("/v1")
	{
		win := v1.Group("/windows")
		{
			win.POST("/sleep",controllers.WindowsController.CallSleep) //休眠计算机
		}
	}



}
