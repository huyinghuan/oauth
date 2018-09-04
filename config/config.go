package config

import (
	"io/ioutil"
	"log"
	"oauth/logger"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

//DbConfig 数据库配置
type DbConfig struct {
	Driver  string
	Connect string
}

type RedisClient struct {
	Addr string `yaml:"addr"`
	Pass string `yaml:"pass"`
	DB   int    `yaml:"db"`
}

type RedisCluster struct {
	URL []string `yaml:"url"`
}

type Redis struct {
	Client RedisClient  `yaml:"client"`
	Cluser RedisCluster `yaml:"cluster"`
}

type Account struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	API  string `yaml:"api"`
}

type Config struct {
	Db      DbConfig `yaml:"db"`
	Dev     bool     `yaml:"dev"`
	Port    string   `yaml:"port"`
	Website string   `yaml:"website"`
	Account Account  `yaml:"account"`
}

var config *Config

//init 读取配置文件
func init() {
	configBytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatalln(err)
	}
	debug := os.Getenv("REGISTRY_AUTH_TOKEN_DEBUG")
	if debug != "" {
		debug = strings.Replace(debug, "\"", "", -1)
		if debug == "false" {
			logger.SetDebug(false)
		}
	}

	admin := os.Getenv("REGISTRY_AUTH_TOKEN_ACCOUNT_USER")
	if admin != "" {
		config.Account.User = strings.Replace(admin, "\"", "", -1)
	}
	pass := os.Getenv("REGISTRY_AUTH_TOKEN_ACCOUNT_PASS")
	if pass != "" {
		config.Account.Pass = strings.Replace(pass, "\"", "", -1)
	}
	api := os.Getenv("REGISTRY_AUTH_TOKEN_ACCOUNT_API")
	if api != "" {
		config.Account.API = strings.Replace(api, "\"", "", -1)
	}
	db := os.Getenv("REGISTRY_AUTH_TOKEN_ACCOUNT_DATABASE")
	if db != "" {
		config.Db.Connect = strings.Replace(db, "\"", "", -1)
	}
}

//Get 获取配置文件
func Get() *Config {
	return config
}
