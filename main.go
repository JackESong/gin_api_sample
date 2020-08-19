package main

import (
	"fmt"
	"net/http"
	"os"
	"gin_api_sample/framework/logger"
	"gin_api_sample/framework/setting"
	"gin_api_sample/project/controller"
	"gin_api_sample/project/entity"
	"gin_api_sample/project/dao"
	"time"
)


func init() {
	fmt.Println("启动初始化方法")
}


func main() {
	fmt.Println("当前时间:" + time.Now().String())
	setting.Setup()
	dao.Setup()

	if len(os.Args) < 1 {
		return
	}
	cfg := &entity.LogConfig{
		Filename:   "catalina.out",
		MaxSize:    50,
		MaxBackups: 50,
		MaxAge:     1800,
	}

	if err := logger.InitLogger(cfg); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	router := controller.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}