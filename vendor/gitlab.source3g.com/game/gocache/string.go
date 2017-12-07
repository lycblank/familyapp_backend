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

func (cache *Cache) Get(key string) (string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()
	return cache.CommandReturnString("GET", key)
}

func (cache *Cache) Set(key string, val string) (ret string, err error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	return cache.CommandReturnString("SET", key, val)
}

func (cache *Cache) SetEx(key string, seconds int64, value interface{}) (ret string, err error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	return cache.CommandReturnString("SETEX", key, seconds, value)
}

func (cache *Cache) MGet(key []interface{}) ([]string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()
	return cache.CommandReturnStringSlice("MGET", key...)
}

func (cache *Cache) SetBit(key string, params ...interface{}) (int, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}
	args = append(args, params...)
	return redis.Int(conn.Do("SETBIT", args...))
}

func (cache *Cache) GetBit(key string, params ...interface{}) (int, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}
	args = append(args, params...)
	return redis.Int(conn.Do("GETBIT", args...))
}
