package main

import (
	"fmt"
	"os"
	"os/signal"
	"status-server/config"
	"status-server/grpc_server"
	"status-server/logging"
	"status-server/redis"
	"syscall"
	"time"
)

func init()  {
		//加载初始化message.toml默认配置
		if err := config.LoadConfigAndSetDefault(); err != nil {
			panic(err.Error())
		}

		//初始化日志配置
		if err := logging.InitZap(&config.GetConf().LogConf); err != nil {
			panic("InitLogger:" + err.Error())
		}

		//初始化redis连接
		redis.InitRedis(config.GetConf().RedisConf)

		//初始化lua脚本
		redis.StatusLua = redis.FromResource(config.GetConf().RedisConf.LuaPath+string(os.PathSeparator)+"status.lua", true)

		//启动grpc服务
		go grpc_server.StartStatusService(config.GetConf())

}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	//测试方法
	//test.Test()

	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("退出", s)
			//服务退出
			grpc_server.StopStatusService()
			redis.Close()
			time.Sleep(time.Second * 3) //等待三秒
			os.Exit(0)
		default:
			fmt.Println("other", s)
		}
	}
}
