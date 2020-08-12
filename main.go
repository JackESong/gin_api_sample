package main

import (
	"fmt"
	"net/http"
	"os"
	"sample_api/framework/conf"
	"sample_api/framework/logger"
	"sample_api/framework/setting"
	"sample_api/project/controller"
)
func main() {
	// load config from conf/conf.json
	if len(os.Args) < 1 {
		return
	}
	// init logger
	cfg := &conf.LogConfig{
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