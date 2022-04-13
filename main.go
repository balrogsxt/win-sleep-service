package main

import (
	"fmt"
	"github.com/kardianos/service"
	"log"
	"math/rand"
	"os"
	"time"
	"xt-service/app"
	"xt-service/services"
)

func main() {
	//初始化日志

	rand.Seed(time.Now().UnixNano())
	if err := app.InitLogger();err != nil {
		log.Fatalln(fmt.Sprintf("初始化日志失败: %s",err.Error()))
	}
	app.Logger.Info("系统初始化成功")

	winSleep := &services.ServiceWindowsSleep{}

	s,err := service.New(winSleep,winSleep.GetConfig())

	if err != nil {
		log.Fatalln(err.Error())
	}
	if 2 > len(os.Args) {
		if err := s.Run();err != nil{
			app.Logger.Fatal("服务启动成功: %s",err.Error())
			return
		}
		app.Logger.Info("服务启动成功")
	}else{
		command := os.Args[1]
		switch command {
		case "install":
			if err := s.Install();err != nil{
				app.Logger.Fatal("安装失败",err.Error())
			}
			app.Logger.Info("安装成功")
			break
		case "uninstall":
			if err := s.Uninstall();err != nil{
				app.Logger.Fatal("卸载失败",err.Error())
			}
			app.Logger.Info("卸载成功")
			break
		case "start":
			status,err := s.Status()
			if err != nil{
				app.Logger.Fatal("启动失败: %s",err.Error())
			}
			app.Logger.Info("启动成功: %d",status)
			break
		case "stop":
			if err := s.Stop();err != nil{
				app.Logger.Fatal("停止失败: %s",err.Error())
			}
			app.Logger.Info("停止成功")
			break
		default:
			app.Logger.Fatal("运行方式不支持")
		}
	}
}