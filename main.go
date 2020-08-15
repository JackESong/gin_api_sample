package main

import (
	"fmt"
	"net/http"
	"os"
	"sample_api/framework/logger"
	"sample_api/framework/setting"
	"sample_api/project/controller"
	"sample_api/project/entity"
	"time"
)


func init() {
	fmt.Println("启动初始化方法")
}


func main() {
	fmt.Println("当前时间:" + time.Now().String())
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
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}