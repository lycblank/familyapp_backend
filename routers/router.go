package routers

import (
	"fmt"
	"gofamily/controllers"
	"gofamily/params"
	"strings"

	familycache "gofamily/cache"

	"github.com/astaxie/beego/context"
	"github.com/garyburd/redigo/redis"
	cache "gitlab.source3g.com/game/gocache"
	"gitlab.source3g.com/game/gorunmanhttp/httpin"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(SetCrossOrigin),
		beego.NSBefore(Auth),
		beego.NSNamespace("/user",
			beego.NSInclude(
				controllers.NewUserController(),
			),
		),
		beego.NSNamespace("/topic",
			beego.NSInclude(
				controllers.NewTopicController(),
			),
		),
	)
	beego.AddNamespace(ns)
}

func Auth(ctx *context.Context) {
	if ctx.Input.IsOptions() {
		// options 方法不进行验证
		return
	}
	if mems := strings.Split(ctx.Input.URL(), "/"); len(mems) >= 3 {
		for i := 2; i < len(mems); i++ {
			if mems[i] == "login" || mems[i] == "register" {
				// 放过登陆与注册操作
				return
			}
		}
	}
	notAuth := params.Response{
		ErrCode: params.NOT_AUTH_ERRCODE,
		ErrMsg:  params.NOT_AUTH_ERRMSG,
	}
	// 从header中获取token
	token := ctx.Input.Header("Authorization")
	if token == "" {
		httpin.JSON(ctx, &notAuth)
		return
	}
	// 验证redis存储的token
	conn := cache.GetRedisConn()
	defer conn.Close()
	authKey := fmt.Sprintf("%s:%s%s", familycache.APP_ID, familycache.REDIS_KEY_USER_AUTH, token)
	if res, err := redis.String(conn.Do("GET", authKey)); err != nil || res == "" {
		fmt.Println(authKey)
		fmt.Println(err, res)
		httpin.JSON(ctx, &notAuth)
		return
	}
}

func SetCrossOrigin(ctx *context.Context) {
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	ctx.Output.Header("Access-Control-Allow-Methods", "*")
	if ctx.Input.IsOptions() {
		httpin.JSON(ctx, params.Response{
			ErrCode: 0,
			ErrMsg:  "success",
		})
	}
}
