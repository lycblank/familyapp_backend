/*
Package cache

Copyright(c) 2000-2016 Quqian Technology Company Limit.
All rights reserved.
创建人：  longyongchun
创建日期：2016-09-18 14:31:34
修改记录：
----------------------------------------------------------------
修改人        |  修改日期    				|    备注
----------------------------------------------------------------
longyongchun  | 2016-09-18 14:31:34		|    创建文件
----------------------------------------------------------------
*/
package cache

import "github.com/garyburd/redigo/redis"

func (cache *Cache) LPush(key string, params ...interface{}) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}
	args = append(args, params...)

	return redis.Int64(conn.Do("LPUSH", args...))
}

func (cache *Cache) RPush(key string, params ...interface{}) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}
	args = append(args, params...)

	return redis.Int64(conn.Do("RPUSH", args...))
}

func (cache *Cache) LPop(key string) (string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}

	return redis.String(conn.Do("LPOP", args...))
}

func (cache *Cache) RPop(key string) (string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}

	return redis.String(conn.Do("RPOP", args...))
}

func (cache *Cache) LRange(key string, start interface{}, stop interface{}) ([]string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, start, stop}
	return cache.CommandReturnStringSlice("LRANGE", args...)
}
