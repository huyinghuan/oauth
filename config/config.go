package config

import (
	"io/ioutil"
	"log"
	"oauth/logger"
	"os"
	"path"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

//DbConfig 数据库配置
type DbConfig struct {
	Driver  string
	Connect string
}

type RedisClient struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"pass"`
	DB       int64  `yaml:"db"`
}

type RedisCluster struct {
	URL []string `yaml:"url"`
}

type Redis struct {
	Client  RedisClient  `yaml:"client"`
	Cluster RedisCluster `yaml:"cluster"`
}

type Account struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	API  string `yaml:"api"`
}

type Config struct {
	Db              DbConfig `yaml:"db"`
	Dev             bool     `yaml:"dev"`
	Port            string   `yaml:"port"`
	Website         string   `yaml:"website"`
	Account         Account  `yaml:"account"`
	Redis           Redis    `yaml:"redis"`
	OpenRegister    bool     `yaml:"open_register"`
	OpenAppRegister bool     `yaml:"open_app_register"`
}

var config *Config

//init 读取配置文件
func init() {
	configPath := "config.yaml"
	if os.Getenv("ProjectPWD") != "" {
		configPath = path.Join(os.Getenv("ProjectPWD"), "config.yaml")
	}
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatalln(err)
	}
	dev := os.Getenv("OPENAUTH_DEV")
	if dev != "" {
		dev = strings.Replace(dev, "\"", "", -1)
		if dev == "false" {
			config.Dev = false
		}
	}

	debug := os.Getenv("OPENAUTH_DEBUG")
	if debug != "" {
		debug = strings.Replace(debug, "\"", "", -1)
		if debug == "false" {
			logger.SetDebug(false)
		}
	}

	admin := os.Getenv("OPENAUTH_ADMIN")
	if admin != "" {
		config.Account.User = strings.Replace(admin, "\"", "", -1)
		config.OpenRegister = false
	}
	pass := os.Getenv("OPENAUTH_ADMIN_PASS")
	if pass != "" {
		config.Account.Pass = strings.Replace(pass, "\"", "", -1)
	}

	db := os.Getenv("OPENAUTH_DATABASE")
	if db != "" {
		config.Db.Connect = strings.Replace(db, "\"", "", -1)
	}

	redisClientAddr := os.Getenv("OPENAUTH_REDIS_CLIENT_ADDR")
	if redisClientAddr != "" {
		config.Redis.Client.Addr = strings.Replace(redisClientAddr, "\"", "", -1)
	}
	redisClientPass := os.Getenv("OPENAUTH_REDIS_CLIENT_PASS")
	if redisClientPass != "" {
		config.Redis.Client.Password = strings.Replace(redisClientPass, "\"", "", -1)
	}
	redisClientDB := os.Getenv("OPENAUTH_REDIS_CLIENT_DB")
	if redisClientDB != "" {
		redisIndex := strings.Replace(redisClientDB, "\"", "", -1)
		config.Redis.Client.DB, _ = strconv.ParseInt(redisIndex, 10, 64)
	}

	openRegister := os.Getenv("OPENAUTH_OPEN_REGISTER")
	if openRegister != "" {
		debug = strings.Replace(openRegister, "\"", "", -1)
		if openRegister == "true" {
			config.OpenRegister = true
		}
	}

}

//Get 获取配置文件
func Get() *Config {
	return config
}
