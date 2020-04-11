package iredis

import (
	"time"
)

func Set(key string, value interface{}) error {
	return client.Set(key, value, 0).Err()
}

func SetEx(key string, value interface{}, seconds time.Duration) error {
	return client.Set(key, value, seconds).Err()
}

func Exist(key string) bool {
	return client.Exists(key).Val() == 1
}

func Del(keys ...string) error {
	for _, key := range keys {
		client.Del(key).Err()
	}
	return nil
}

func Get(key string) (string, error) {
	return client.Get(key).Result()
}

func RPush(key string, values ...string) error {
	return client.RPush(key, values).Err()
}

func LRange(key string, start int64, end int64) ([]string, error) {
	return client.LRange(key, start, end).Result()
}
