package main

import (
	"Go-details/global"
	"Go-details/internal/model"
	"Go-details/internal/routers"
	"Go-details/pkg/logger"
	"Go-details/pkg/setting"
	"Go-details/pkg/tracer"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

/**
* @Author: super
* @Date: 2020-09-29 14:50
* @Description:
**/

var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

func init() {
	//读取命令行参数
	err := setupFlag()
	if err != nil {
		log.Printf("init.setupFlag err: %v\n", err)
	}
	//初始化配置
	err = setupSetting()
	if err != nil {
		log.Printf("init setupSetting err: %v\n", err)
	}
	//初始化日志
	err = setupLogger()
	if err != nil {
		log.Printf("init setupLogger err: %v\n", err)
	}
	//初始化数据库
	err = setupDBEngine()
	if err != nil {
		log.Printf("init setupDBEngine err: %v\n", err)
	}
	//初始化追踪
	err = setupTracer()
	if err != nil {
		global.Logger.Fatalf(context.Background(), "init.setupTracer err: %v", err)
	}
}

// @title 技术细节
// @version 1.0
// @description 注重技术细节
// @Github https://github.com/golang-collection/Go-details
func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := pingServer(); err != nil {
			global.Logger.Errorf(context.Background(), "The server has no response, or it might took too long to start up.")
		}
		global.Logger.Info(context.Background(), "The server has been deployed successfully.")
	}()

	global.Logger.Infof(context.Background(), "Start to listening the incoming requests on http address :%s", global.ServerSetting.HttpPort)
	err := s.ListenAndServe()
	if err != nil {
		global.Logger.Fatalf(context.Background(), "start listen server err: %v", err)
	}
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(":" + global.ServerSetting.HttpPort + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		global.Logger.Info(context.Background(), "Waiting for the server, retry in 1 second.")
	}
	return errors.New("cannot connect to the server")
}

func setupSetting() error {
	newSetting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	fmt.Println("log file name ", fileName)
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("bedtimeStory", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}