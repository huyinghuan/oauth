package iredis

import (
	"fmt"
)

type appCache struct {
	Format string
}

var AppCache = appCache{Format: "app:%s"}

func (c *appCache) SetAll(clientID string, pk string, cb string) error {
	key := fmt.Sprintf(c.Format, clientID)
	return client.HMSet(key, "callback", cb, "pk", pk).Err()
}

func (c *appCache) SetCallback(clientID string, cb string) error {
	key := fmt.Sprintf(c.Format, clientID)
	return client.HSet(key, "callback", cb).Err()
}

func (c *appCache) GetCallback(clientID string) (string, error) {
	key := fmt.Sprintf(c.Format, clientID)
	return client.HGet(key, "callback").Result()
}

func (c *appCache) SetPrivateKey(clientID string, pk string) error {
	key := fmt.Sprintf(c.Format, clientID)
	return client.HSet(key, "pk", pk).Err()
}

func (c *appCache) GetPrivateKey(clientID string) (string, error) {
	key := fmt.Sprintf(c.Format, clientID)
	return client.HGet(key, "pk").Result()
}

func (c *appCache) SetMode(clientID string, mode string) error {
	key := fmt.Sprintf(c.Format, clientID)
	return client.HSet(key, "mode", mode).Err()
}

func (c *appCache) GetMode(clientID string) (string, error) {
	key := fmt.Sprintf(c.Format, clientID)
	return client.HGet(key, "mode").Result()
}

func (c *appCache) Clear(clientID string) error {
	key := fmt.Sprintf(c.Format, clientID)
	return Del(key)
}

func (c *appCache) Exist(clientID string) bool {
	key := fmt.Sprintf(c.Format, clientID)
	return Exist(key)
}
