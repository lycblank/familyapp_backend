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

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisConfig struct {
	Network        string `json:"network"`
	Address        string `json:"address"`
	Password       string `json:"password"`
	ConnectTimeout int    `json:"connecttimeout"`
	ReadTimeout    int    `json:"readtimeout"`
	WriteTimeout   int    `json:"writetimeout"`
	MaxActive      int    `json:"maxactive"`
	MaxIdle        int    `json:"maxidle"`
	IdleTimeout    int    `json:"idletimeout"`
	Wait           bool   `json:"wait"`
	DB             string `json:"db"`
}

type Cache struct {
	redisPool *redis.Pool
}

var DefaultCache Cache

func InitRedisPool(configs ...RedisConfig) {
	DefaultCache.RedisPool(configs...)
}

// 二次封装
func GetRedisConn() redis.Conn {
	return DefaultCache.GetRedisConn()
}

func GetRedisPool() *redis.Pool {
	return DefaultCache.redisPool
}

func (cache *Cache) RedisPool(configs ...RedisConfig) *redis.Pool {
	if cache.redisPool == nil {
		if len(configs) == 0 {
			return nil
		}

		cache.newRedisPool(configs[0])
	}
	return cache.redisPool
}

func (cache *Cache) newRedisPool(config RedisConfig) {
	cache.redisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {

			var connect_timeout time.Duration = time.Duration(config.ConnectTimeout) * time.Second
			var read_timeout = time.Duration(config.ReadTimeout) * time.Second
			var write_timeout = time.Duration(config.WriteTimeout) * time.Second

			c, err := redis.DialTimeout(config.Network, config.Address, connect_timeout, read_timeout, write_timeout)

			if err != nil {
				return nil, err
			}
			if len(config.Password) > 0 {
				if _, err := c.Do("AUTH", config.Password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if len(config.DB) > 0 {
				if _, err = c.Do("SELECT", config.DB); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: time.Duration(config.IdleTimeout) * time.Second,
		Wait:        config.Wait,
	}
}

func (cache *Cache) CommandReturnStringSlice(commandName string, args ...interface{}) ([]string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	resp, err := redis.Strings(conn.Do(commandName, args...))
	if err != nil {
		if resp == nil || len(resp) <= 0 {
			return nil, redis.ErrNil
		}
	}
	return resp, err
}

func (cache *Cache) CommandReturnString(commandName string, args ...interface{}) (string, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	resp, err := redis.String(conn.Do(commandName, args...))
	if err != nil {
		if resp == "" {
			return "", redis.ErrNil
		}
	}
	return resp, err
}

func (cache *Cache) CommandReturnInterface(commandName string, args ...interface{}) (interface{}, error) {
	conn := cache.RedisPool().Get()
	defer conn.Close()

	return conn.Do(commandName, args...)
}

func (cache *Cache) GetRedisConn() redis.Conn {
	return cache.RedisPool().Get()
}
