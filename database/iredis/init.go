package iredis

import (
	"shorturl/config"

	"log"

	redis "gopkg.in/redis.v3"
)

var client *redis.ClusterClient
var
func init() {
	conf := config.Get()

	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: conf.Redis.Url,
	})
	if err := client.Ping().Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("redis 连接成功")
	}
}
