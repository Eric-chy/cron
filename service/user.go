package service

import (
	"cron/global"
	"cron/model"
	"cron/pkg/httpclient"
	"fmt"
	"sync"
	"time"
)

func (svc *Service) GetList(name []interface{}) {
	user := model.User{}
	if len(name) > 0 {
		user.Name = name[0].(string)
	}
	page := 1
	pageSize := 10
	pageOffset := (page - 1) * pageSize
	hotels, _ := user.List(svc.engine, pageOffset, pageSize)
	fmt.Println(hotels)
}

func (svc *Service) Curl() {
	res, err := httpclient.New().Timeout(3 * time.Second).Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Close()
	s := res.ReadAllString()
	fmt.Println(s)
	//r := gjson.JsonDecode(s)
}

func (svc *Service) MultiThirdApi(num []interface{}) {
	var wg sync.WaitGroup
	var n int
	if len(num) > 0 {
		n = num[0].(int)
	}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					global.Logger.Error(err)
				}
			}()
			defer wg.Done()
			res, err := httpclient.New().Timeout(3 * time.Second).Get("http://www.baidu.com")
			if err != nil {
				fmt.Println(err)
			}
			defer res.Close()
			s := res.ReadAllString()
			fmt.Println(s)
		}()
	}
	wg.Wait()
}
