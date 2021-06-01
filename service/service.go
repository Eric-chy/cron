package service

import (
	"cron/global"
	"github.com/jinzhu/gorm"
	"reflect"
)

type Service struct {
	engine *gorm.DB
}

func NewService() *Service {
	return &Service{engine: global.DBEngine}
}

func (svc *Service) HandleFunc(funcName string, args ...interface{}) func() {
	return func() {
		getValue := reflect.ValueOf(svc)
		getType := reflect.TypeOf(svc)

		met, ok := getType.MethodByName(funcName)
		if ok {
			handler := getValue.MethodByName(funcName)
			if len(args) > 0 {
				handler.Call([]reflect.Value{reflect.ValueOf(args)})
			} else {
				met.Func.Call([]reflect.Value{reflect.ValueOf(svc)})
			}
		} else {
			global.Logger.Error("func " + funcName + " not exists.")
		}
	}
}
