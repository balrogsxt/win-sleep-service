package controllers

import (
	"github.com/gin-gonic/gin"
	"os/exec"
	"time"
	"xt-service/app"
)

var WindowsController = new(_WindowsController)

type _WindowsController struct {
}
// CallSleep 休眠计算机
func (_WindowsController) CallSleep(c *gin.Context) {
	go func() {
		time.Sleep(time.Second) //延迟1秒处理,不然数据还没返回就休眠了
		res,err := exec.Command("rundll32.exe","powrprof.dll,SetSuspendState","0,1,0").Output()
		if err != nil {
			app.Logger.Warn("休眠计算机失败: %s",err.Error())
		}else{
			app.Logger.Info("休眠计算机响应成功: %s",res)
		}
	}()
	app.Response.Ok(c,nil,"请求已提交,等待目标计算机响应")
}