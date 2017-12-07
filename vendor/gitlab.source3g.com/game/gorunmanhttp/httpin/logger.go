/*
 Package xxx

 Copyright(c) 2000-2016 Quqian Technology Company Limit.
 All rights reserved.
 创建人：  longyongchun
 创建日期：2016-09-13 12:51:37
 修改记录：
 ----------------------------------------------------------------
 修改人        |  修改日期    				|    备注
 ----------------------------------------------------------------
 longyongchun  | 2016-09-13 12:51:37		|    创建文件
 ----------------------------------------------------------------
*/

package httpin

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/logs"
)

var log *logs.BeeLogger

func SetLogger(loger *logs.BeeLogger) {
	log = loger
}

func Info(v interface{}, format string, args ...interface{}) {
	logString := GetLogIDString(v)
	if log != nil {
		log.Info(logString+format, args...)
	}
}

func Error(v interface{}, format string, args ...interface{}) {
	logString := GetLogIDString(v)
	if log != nil {
		log.Error(logString+format, args...)
	}
}

func Debug(v interface{}, format string, args ...interface{}) {
	logString := GetLogIDString(v)
	if log != nil {
		log.Debug(logString+format, args...)
	}
}

func Warning(v interface{}, format string, args ...interface{}) {
	logString := GetLogIDString(v)
	if log != nil {
		log.Warning(logString+format, args...)
	}
}

func GetLogIDString(arg interface{}) string {
	v := reflect.ValueOf(arg)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct {
		numField := v.NumField()
		for i := 0; i < numField; i++ {
			if v.Type().Field(i).Name == "LogID" && v.Type().Field(i).Type.Kind() == reflect.String {
				return fmt.Sprintf("log id:[%s] ", v.Field(i).String())
			}
		}
	}
	return ""
}
