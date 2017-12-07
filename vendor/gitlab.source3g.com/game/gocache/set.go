package cache

import "github.com/garyburd/redigo/redis"

func (cache *Cache) Scard(key string) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	return redis.Int64(conn.Do("SCARD", key))
}
