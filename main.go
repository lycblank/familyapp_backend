package main

import (
	"bytes"
	"fmt"
	_ "gofamily/routers"
	"os"

	"gofamily/httpin"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	cache "gitlab.source3g.com/game/gocache"
	config "gitlab.source3g.com/game/goconfig"
)

var logger = logs.NewLogger(10000)

func initLogger(runmode string) {
	logger.SetLogger(logs.AdapterFile, `{"filename":"logs/family.log", "maxdays":30}`)
	logger.Async(10000)
	logger.EnableFuncCallDepth(true)
	// 整个项目中的日志处理都进行封装了一层
	logger.SetLogFuncCallDepth(logger.GetLogFuncCallDepth() + 1)
	if runmode == "prod" {
		logger.SetLevel(logs.LevelInformational)
	} else {
		logger.SetLevel(logs.LevelDebug)
	}

	httpin.SetLogger(logger)
}

func initMysql(runmode string) {
	var buf bytes.Buffer
	buf.WriteString(config.Conf.Mysql.UserName)
	buf.WriteByte(':')
	buf.WriteString(config.Conf.Mysql.Password)
	buf.WriteString(`@tcp(`)
	buf.WriteString(config.Conf.Mysql.Host)
	buf.WriteByte(':')
	buf.WriteString(config.Conf.Mysql.Port)
	buf.WriteString(`)/`)
	buf.WriteString(config.Conf.Mysql.DBName)
	buf.WriteString(`?charset=`)
	buf.WriteString(config.Conf.Mysql.Charset)
	if err := orm.RegisterDataBase("default", "mysql", buf.String()); err != nil {
		fmt.Printf("register mysql failed. error:%s\n", err)
		logger.Error("register mysql failed. error:%s", err)
		os.Exit(1)
	}
	orm.SetMaxIdleConns("default", config.Conf.Mysql.MaxIdleConns)
	orm.SetMaxOpenConns("default", config.Conf.Mysql.MaxOpenConns)
	if runmode == "dev" {
		orm.Debug = true
	}
}

func main() {
	runmode := os.Getenv("runmode")
	if runmode == "" {
		runmode = "dev"
	}
	cfg := config.Init(runmode)
	cache.InitRedisPool(cfg.Redis)
	initLogger(runmode)
	initMysql(runmode)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
