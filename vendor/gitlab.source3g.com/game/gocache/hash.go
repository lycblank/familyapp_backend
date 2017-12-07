/*
Package cache

Copyright(c) 2000-2016 Quqian Technology Company Limit.
All rights reserved.
创建人：  longyongchun
创建日期：2016-09-18 10:52:43
修改记录：
----------------------------------------------------------------
修改人        |  修改日期    				|    备注
----------------------------------------------------------------
longyongchun  | 2016-09-18 10:52:43		|    创建文件
----------------------------------------------------------------
*/
package cache

import "github.com/garyburd/redigo/redis"

func (cache *Cache) HMGet(key string, fields ...string) ([]string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}
	for _, val := range fields {
		args = append(args, val)
	}
	return cache.CommandReturnStringSlice("HMGET", args...)
}

func (cache *Cache) HGet(key string, field string) (string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, field}

	return cache.CommandReturnString("HGET", args...)

}

func (cache *Cache) HMSet(key string, values ...interface{}) (string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}
	args = append(args, values...)

	return cache.CommandReturnString("HMSET", args...)
}

func (cache *Cache) HSet(key string, field string, value interface{}) (bool, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, field, value}

	return redis.Bool(conn.Do("HSET", args...))
}

func (cache *Cache) HIncrBy(key string, filed interface{}, val interface{}) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, filed, val}

	return redis.Int64(conn.Do("HINCRBY", args...))
}
