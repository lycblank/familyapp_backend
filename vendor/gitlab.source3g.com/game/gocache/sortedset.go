/*
Package xxx

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

func (cache *Cache) ZRevRange(key string, start int64, stop int64, params ...interface{}) ([]string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, start, stop}
	args = append(args, params...)
	return cache.CommandReturnStringSlice("ZREVRANGE", args...)
}

func (cache *Cache) ZRevRangeByScore(key string, start interface{}, stop interface{}, params ...interface{}) ([]string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, start, stop}
	args = append(args, params...)
	return cache.CommandReturnStringSlice("ZREVRANGEBYSCORE", args...)
}

func (cache *Cache) ZAdd(key string, params ...interface{}) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}
	args = append(args, params...)

	return redis.Int64(conn.Do("ZADD", args...))
}
func (cache *Cache) ZCount(key string, min interface{}, max interface{}) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, min, max}

	return redis.Int64(conn.Do("ZCOUNT", args...))
}

func (cache *Cache) ZScore(key string, member interface{}) (string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, member}

	return redis.String(conn.Do("ZSCORE", args...))
}

func (cache *Cache) ZScoreAsInt64(key string, member interface{}) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, member}

	return redis.Int64(conn.Do("ZSCORE", args...))
}

func (cache *Cache) ZRange(key string, params ...interface{}) ([]string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}
	args = append(args, params...)
	return cache.CommandReturnStringSlice("ZRANGE", args...)
}

func (cache *Cache) ZIncrBy(key string, increment int64, member interface{}) (string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, increment, member}

	return cache.CommandReturnString("ZINCRBY", args...)
}

func (cache *Cache) ZRem(key string, members ...interface{}) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key}
	args = append(args, members...)

	return redis.Int64(conn.Do("ZREM", args...))
}

func (cache *Cache) ZUnionStore(dest string, params ...interface{}) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{dest}
	args = append(args, params...)

	return redis.Int64(conn.Do("ZUNIONSTORE", args...))
}

func (cache *Cache) Zcard(key string) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	return redis.Int64(conn.Do("ZCARD", key))
}

func (cache *Cache) ZRemRangeByScore(key string, min interface{}, max interface{}) (int64, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	args := []interface{}{key, min, max}

	return redis.Int64(conn.Do("ZREMRANGEBYSCORE", args...))
}
