package main

import (
	"blog/conf"
	"blog/model"
	"blog/router"
	"os"
	"os/signal"
	"syscall"
	"github.com/zxysilent/logs"
)

// @Title Blog’s Api文档
// @Version 1.0
// @Description token传递方式包括 [get/post]token 、[header] Authorization=Bearer xxxx
// @Description 数据传递方式包括 json、formData 推荐使用 json
// @Description /api/* 公共访问
// @Description /adm/* 必须传入 token
// @Host 127.0.0.1:88
// @BasePath /
func main() {
	// 打印日志
	logs.Info("app initializing")
	// 初始化配置
	conf.Init()
	// 初始化模型
	model.Init()

	// 创建channel，用于监听系统信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	logs.Info("app running")

	// 启用协程，用于运行项目
	go router.RunApp()

	// 登录退出信号
	<-quit

	logs.Info("app quitted")

	// 释放日志
	logs.Flush()
}
