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

func (cache *Cache) Expire(key string, timeout int) error {
	conn := cache.RedisPool().Get()
	defer conn.Close()
	_, err := conn.Do("EXPIRE", key, timeout)
	return err
}

func (cache *Cache) Exists(key string) (bool, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}

	return redis.Bool(conn.Do("EXISTS", args...))
}

func (cache *Cache) Del(keys ...string) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{}
	for _, key := range keys {
		args = append(args, key)
	}

	return redis.Int64(conn.Do("DEL", args...))
}

func (cache *Cache) IncrBy(key string, num interface{}) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, num}

	return redis.Int64(conn.Do("INCRBY", args...))
}

func (cache *Cache) Rename(oldKey string, newKey string) (string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()
	args := []interface{}{oldKey, newKey}
	return cache.CommandReturnString("RENAME", args...)
}
