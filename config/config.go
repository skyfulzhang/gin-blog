package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"time"
)

// 应用信息
type APPConfig struct {
	Desc       string `yaml:"desc" json:"desc"`
	Addr       string `yaml:"addr" json:"addr"`
	ConfigFile string `yaml:"configFile" json:"configFile"`
	Version    string `yaml:"version" json:"version"`
	Env        string `yaml:"env" json:"env"`
}

// MySQL信息
type MysqlConfig struct {
	Host      string `yaml:"host" json:"host"`
	Port      string `yaml:"port" json:"port"`
	Username  string `yaml:"username" json:"username"`
	Password  string `yaml:"password" json:"password"`
	Database  string `yaml:"database" json:"database"`
	Charset   string `yaml:"charset" json:"charset"`
	Parsetime string `yaml:"parsetime" json:"parsetime"`
}

// Redis
type RedisConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Default  string `yaml:"default" json:"default"`
	Timeout  string `yaml:"timeout" json:"timeout"`
}
type JWTConfig struct {
	Secret string        `yaml:"secret" json:"secret"`
	Issuer string        `yaml:"issuer" json:"issuer"`
	Expire time.Duration `yaml:"expire" json:"expire"`
}

// ServerConfig 配置信息
type ServerConfig struct {
	App   APPConfig   `yaml:"app" json:"app"`
	Mysql MysqlConfig `yaml:"mysql" json:"mysql"`
	Redis RedisConfig `yaml:"redis" json:"redis"`
	JWT   JWTConfig   `yaml:"jwt" json:"jwt"`
}

func InitConfig() {
	var config = ServerConfig{}
	var err error
	dir, _ := os.Getwd()
	fmt.Println("dir=", dir)
	viper.SetConfigFile(dir + "/controller/app.yaml")
	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed:%s\n", err))
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("unmarshal failed:%s\n", err))
	}
	fmt.Println("config:", config)
	// 动态监测配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变")
		if err := viper.Unmarshal(&config); err != nil {
			panic(fmt.Errorf("配置重载失败:%s\n", err))
		}
	})
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("配置重载失败:%s\n", err))
	}
}
