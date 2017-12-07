package httpin

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func JSON(ctx *context.Context, param interface{}, logID ...interface{}) {
	if ctx == nil {
		return
	}
	ctx.Output.JSON(param, false, false)

	if beego.BConfig.RunMode == "dev" {
		// 打印debug日志
		datas, _ := json.Marshal(param)
		var logid interface{}
		if logID != nil && len(logID) > 0 {
			logid = logID[0]
		}
		Debug(logid, "response:%s", string(datas))
	}

}
