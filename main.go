package main

import (
	"cron/config"
	"cron/global"
	"cron/pkg/logger"
	"cron/service"
	"fmt"
	"github.com/robfig/cron"
)

func init() {
	config.Init()
	logger.Init()
	//model.Init()
	//gredis.Init()
	//mgodb.Init()
}
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			global.Logger.Error(err)
		}
	}()
	svc := service.NewService()
	c := cron.New()
	//spec := "*/2 * * * * ?"     //every 2s
	//spec := "0 5 10 * * ?"      //10：05：00
	//spec := "* * 10 * * ?"      //10
	//spec := "* 1-59/10 * * * ?" //1-59m  10m
	/*
		┌─────────────second 范围 (0 - 60)
		│ ┌───────────── min (0 - 59)
		│ │ ┌────────────── hour (0 - 23)
		│ │ │ ┌─────────────── day of month (1 - 31)
		│ │ │ │ ┌──────────────── month (1 - 12)
		│ │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
		│ │ │ │ │ │                  Saturday)
		│ │ │ │ │ │
		│ │ │ │ │ │
		*  *  *  *  *  *
	*/
	//c.AddFunc("@every 5s", svc.HandleFunc("GetList", "zhangsan"))
	//c.AddFunc("@every 10s", svc.HandleFunc("Curl"))
	c.AddFunc("@every 2s", svc.HandleFunc("MultiThirdApi", 10))
	c.Start()
	select {}
}
