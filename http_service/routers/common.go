package routers

import (
	"github.com/gin-gonic/gin"
	"xt-service/http_service/middlewares"
)

func RegisterRouter(g *gin.Engine)  {
	//注册中间件
	g.Use(
		middlewares.Recover(),
		middlewares.Cors(),
	)
	registerApi(g)
}
