package config

import (
	"github.com/spf13/viper"
	"time"
)

var Conf Yaml

type Yaml struct {
	App struct {
		Env             string        `yaml:"env"`
		Version         string        `yaml:"version"`
		ReadTimeout     time.Duration `yaml:"readTimeout"`
		WriteTimeout    time.Duration `yaml:"writeTimeout"`
		AppName         string        `yaml:"appName"`
		LogDir          string        `yaml:"logDir"`
		MaxPageSize     int           `yaml:"maxPageSize"`
		DefaultPageSize int           `yaml:"defaultPageSize"`
	}
	Database struct {
		Driver      string        `yaml:"driver"`
		Protocol    string        `yaml:"tcp"`
		Host        string        `yaml:"host"`
		Port        int           `yaml:"port"`
		User        string        `yaml:"user"`
		Password    string        `yaml:"password"`
		Name        string        `yaml:"name"`
		Prefix      string        `yaml:"prefix"`
		RunMode     string        `yaml:"runMode"`
		MaxIdles    int           `yaml:"maxIdles"`
		MaxOpens    int           `yaml:"maxOpens"`
		MaxLifetime time.Duration `yaml:"maxLifetime"`
	}
	Redis struct {
		Driver      string        `yaml:"driver"`
		Protocol    string        `yaml:"tcp"`
		Host        string        `yaml:"host"`
		Port        string        `yaml:"port"`
		Password    string        `yaml:"password"`
		MaxIdle     int           `yaml:"maxIdle"`
		MaxActive   int           `yaml:"maxActive"`
		IdleTimeout time.Duration `yaml:"IdleTimeout"`
	}
	Sentry struct {
		Dsn string `yaml:"dsn"`
	}
	Email struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		UserName string `yaml:"userName"`
		Password string `yaml:"password"`
		IsSSL    bool   `yaml:"isSSL"`
		From     string `yaml:"from"`
	}
	Es struct {
		Link                string `yaml:"link"`
		MaxIdleConnsPerHost int    `yaml:"maxIdleConnsPerHost"`
	}
	Mongo struct {
		Host        string        `yaml:"host"`
		MaxPoolSize uint64        `yaml:"maxPoolSize"`
		Timeout     time.Duration `yaml:"timeout"`
	}
}

var vp *viper.Viper

func Init() {
	vp = viper.New()
	vp.SetConfigName("dev")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//直接整个解析
	err = vp.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
