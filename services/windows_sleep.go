package services

import (
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"net/http"
	"xt-service/app"
	"xt-service/http_service/routers"
)

type ServiceWindowsSleep struct {
	httpServer *http.Server
}

func (p *ServiceWindowsSleep) GetConfig() *service.Config {
	return &service.Config{
		Name: "Windows-Sleep",
		DisplayName: "Windows-Sleep",
		Description: "提供HTTP API休眠当前计算机",
	}
}
func (p *ServiceWindowsSleep) Start(s service.Service) error {
	//启动Http服务器
	router := gin.New()
	routers.RegisterRouter(router)
	go func() {
		if err := router.Run(":10269");err != nil {
			app.Logger.Error("启动Http服务器失败: %s",err.Error())
			s.Stop()
			app.Logger.Fatal("终止运行")
		}
	}()
	return nil
}
func (p *ServiceWindowsSleep) Stop(s service.Service) error {
	//if err := p.httpServer.Shutdown(context.Background());err != nil {
	//	app.Logger.Error("%s - 停止失败: %s",p.GetConfig().Name,err.Error())
	//	return err
	//}
	app.Logger.Info("%s - 已停止",p.GetConfig().Name)
	return nil
}

