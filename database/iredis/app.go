package iredis

import (
	"fmt"
)

type appCache struct {
	Format    string
	MapFormat string
}

var AppCache = appCache{Format: "app:%d", MapFormat: "app:map:%s"}

func (c *appCache) SetMap(id int64, clientID string) error {
	key := fmt.Sprintf(c.MapFormat, clientID)
	return client.Set(key, id, 0).Err()
}

func (c *appCache) GetMap(clientID string) (int64, error) {
	key := fmt.Sprintf(c.MapFormat, clientID)
	return client.Get(key).Int64()
}

func (c *appCache) SetAll(appID int64, pk string, cb string, mode string) error {
	key := fmt.Sprintf(c.Format, appID)
	return client.HMSet(key, "callback", cb, "pk", pk, "mode", mode).Err()
}

func (c *appCache) SetCallback(appID int64, cb string) error {
	key := fmt.Sprintf(c.Format, appID)
	return client.HSet(key, "callback", cb).Err()
}

func (c *appCache) GetCallback(appID int64) (string, error) {
	key := fmt.Sprintf(c.Format, appID)
	return client.HGet(key, "callback").Result()
}

func (c *appCache) SetPrivateKey(appID int64, pk string) error {
	key := fmt.Sprintf(c.Format, appID)
	return client.HSet(key, "pk", pk).Err()
}

func (c *appCache) GetPrivateKey(appID int64) (string, error) {
	key := fmt.Sprintf(c.Format, appID)
	return client.HGet(key, "pk").Result()
}

func (c *appCache) SetMode(appID int64, mode string) error {
	key := fmt.Sprintf(c.Format, appID)
	return client.HSet(key, "mode", mode).Err()
}

func (c *appCache) GetMode(appID int64) (string, error) {
	key := fmt.Sprintf(c.Format, appID)
	return client.HGet(key, "mode").Result()
}

func (c *appCache) Clear(appID int64) error {
	key := fmt.Sprintf(c.Format, appID)
	return Del(key)
}

func (c *appCache) Exist(clientID string) bool {
	key := fmt.Sprintf(c.MapFormat, clientID)
	return Exist(key)
}
