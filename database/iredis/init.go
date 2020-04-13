package iredis

import (
	"oauth/config"

	"log"

	redis "github.com/go-redis/redis/v7"
)

var cluster *redis.ClusterClient
var client *redis.Client

func Init() {
	conf := config.Get()

	// client = redis.NewClusterClient(&redis.ClusterOptions{
	// 	Addrs: conf.Redis.Cluster.Url,
	// })
	client = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Client.Addr,
		Password: conf.Redis.Client.Password,
		DB:       conf.Redis.Client.DB,
	})
	if err := client.Ping().Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("redis 连接成功")
	}
}
