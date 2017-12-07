package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	cache "gitlab.source3g.com/game/gocache"
)

var Conf *Config

type ServerConfig struct {
	Host              string `json:"host"`
	MsgType           int32  `json:"msg_type" description:"1: 标识json 2:标识msgpack"`
	Name              string `json:"name" description:"进程名"`
	ActEndTime        int    `json:"act_end_time" description:"活动结束时间"`
	ActTotalTime      int    `json:"act_total_time" description:"活动总时间"`
	FinalRankWaitTime int    `json:"final_rank_wait_time" description:"最终下发排名需要等待的时间"`
}

type MysqlConfig struct {
	UserName     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	DBName       string `json:"db_name"`
	Charset      string `json:"charset"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

type Config struct {
	Redis  cache.RedisConfig `json:"redis"`
	Server ServerConfig      `json:"server"`
	Mongo  MongoConfig       `json:"mongodb"`
	Mysql  MysqlConfig       `json:"mysql"`
}

type MongoConfig struct {
	Host string `json:"host"`
}

func Init(runmode string) *Config {
	confPath := "conf/dev.json"
	if runmode == "prod" {
		confPath = "conf/prod.json"
	}
	datas, err := ioutil.ReadFile(confPath)
	if err != nil {
		fmt.Printf("read file failed. error:%s", err)
		os.Exit(1)
	}
	Conf = &Config{}
	if err := json.Unmarshal(datas, Conf); err != nil {
		fmt.Printf("unmarshal data failed. error:%s", err)
		os.Exit(1)
	}
	return Conf
}
