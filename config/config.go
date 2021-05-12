package config

import (
	"embed"
	"log"
	"os"

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
	DB       int    `yaml:"db"`
}

type RedisCluster struct {
	URL []string `yaml:"url"`
}

type Redis struct {
	Client  RedisClient  `yaml:"client"`
	Cluster RedisCluster `yaml:"cluster"`
}

// 默认管理员
type Account struct {
	User           string `yaml:"user"`
	Pass           string `yaml:"pass"`
	ResetOnRestart bool   `yaml:"resetOnRestart"`
}

func (a Account) IsAdmin(username string) bool {
	return a.User == username
}

type Config struct {
	Db               DbConfig `yaml:"db"`
	Dev              bool     `yaml:"dev"`
	Website          string   `yaml:"website"`
	Account          Account  `yaml:"account"`
	Redis            Redis    `yaml:"redis"`
	OpenRegister     bool     `yaml:"open_register"`
	OpenAppRegister  bool     `yaml:"open_app_register"`
	RedisCacheFromDB bool     `yaml:"redis_cache_from_db"`
	JWTSecret        string   `yaml:"jwt"`
}

var config *Config

//go:embed *
var configFile embed.FS
var RunMode = ""

// init 读取配置文件
func init() {
	runmode := os.Getenv("RUN_MODE")
	if runmode == "" {
		runmode = RunMode
	}
	configPath := "config.dev.yml"
	if runmode == "product" {
		configPath = "config.yml"
	}
	log.Println("当前配置:", configPath)
	body, err := configFile.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(body, &config)
	if err != nil {
		log.Fatalln(err)
	}
}

//Get 获取配置文件
func Get() *Config {
	return config
}
