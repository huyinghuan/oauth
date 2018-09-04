package iredis

func Set(key string, value interface{}) error {
	return client.Set(key, value, 0).Err()
}

func Del(keys ...string) error {
	var queue []string
	for _, key := range keys {
		if client.Exists(key).Val() {
			queue = append(queue, key)
		}
	}
	if len(queue) == 0 {
		return nil
	}
	return client.Del(queue...).Err()
}

func Get(key string) (string, error) {
	return client.Get(key).Result()
}

func RPush(key string, values ...string) error {
	return client.RPush(key, values...).Err()
}

func LRange(key string, start int64, end int64) ([]string, error) {
	return client.LRange(key, start, end).Result()
}
