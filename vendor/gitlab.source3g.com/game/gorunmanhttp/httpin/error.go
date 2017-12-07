package httpin

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 错误时返回
type GeneralError struct {
	HttpCode uint32 `json:"http_code" description:"http返回码"`
	ErrCode  uint32 `json:"err_code" description:"服务器具体错误码"`
	ErrMsg   string `json:"err_msg" description:"服务器错误描述"`
}

func (g *GeneralError) Error() string {
	datas, _ := json.Marshal(g)
	return string(datas)
}

func SystemError(httpCode uint32, errMsg ...string) error {
	errMessage := "system error: "
	if errMsg != nil && len(errMsg) > 0 {
		errMessage += strings.Join(errMsg, ",")
	}
	return &GeneralError{
		HttpCode: httpCode,
		ErrCode:  3,
		ErrMsg:   errMessage,
	}
}

func GetSystemError(httpCode uint32, errMsg ...string) error {
	errMessage := "system error: "
	if errMsg != nil && len(errMsg) > 0 {
		errMessage += strings.Join(errMsg, ",")
	}
	return &GeneralError{
		HttpCode: httpCode,
		ErrCode:  3,
		ErrMsg:   errMessage,
	}
}

func BadParamer(httpCode uint32, errMsg ...string) error {
	errMessage := "bad paramer: "
	if errMsg != nil && len(errMsg) > 0 {
		errMessage += strings.Join(errMsg, ",")
	}
	return &GeneralError{
		HttpCode: httpCode,
		ErrCode:  2,
		ErrMsg:   errMessage,
	}
}

func ConflictError(httpCode uint32, errMsg ...string) {
	errMessage := "conflict error: "
	if errMsg != nil && len(errMsg) > 0 {
		errMessage += strings.Join(errMsg, ",")
	}
	panic(&GeneralError{
		HttpCode: httpCode,
		ErrCode:  16,
		ErrMsg:   errMessage,
	})
}

func GetConflictError(httpCode uint32, errMsg ...string) error {
	errMessage := "conflict error: "
	if errMsg != nil && len(errMsg) > 0 {
		errMessage += strings.Join(errMsg, ",")
	}
	return &GeneralError{
		HttpCode: httpCode,
		ErrCode:  16,
		ErrMsg:   errMessage,
	}
}

func SourceDeleted(httpCode uint32, errMsg ...string) {
	errMessage := "source delete:"
	if errMsg != nil && len(errMsg) > 0 {
		errMessage += strings.Join(errMsg, ",")
	}
	panic(&GeneralError{
		HttpCode: httpCode,
		ErrCode:  20,
		ErrMsg:   errMessage,
	})
}

func GetSourceDeletedError(httpCode uint32, errMsg ...string) error {
	errMessage := "source delete:"
	if errMsg != nil && len(errMsg) > 0 {
		errMessage += strings.Join(errMsg, ",")
	}
	/*datas, _ := json.Marshal(GeneralError{
		HttpCode: httpCode,
		ErrCode:  20,
		ErrMsg:   errMessage,
	})*/
	return &GeneralError{
		HttpCode: httpCode,
		ErrCode:  20,
		ErrMsg:   errMessage,
	}
	//return core.ServerError(string(datas))
}

func NoData(httpCode uint32, errMsg ...string) {
	errMessage := "no data:"
	if errMsg != nil && len(errMsg) > 0 {
		errMessage += strings.Join(errMsg, ",")
	}
	panic(&GeneralError{
		HttpCode: httpCode,
		ErrCode:  22,
		ErrMsg:   errMessage,
	})
}

func InvalidSid(httpCode uint32, errMsgs ...string) {
	errMsg := "invalid sid"
	if len(errMsgs) > 0 {
		errMsg = errMsg + ":"
		for idx, msg := range errMsgs {
			errMsg = fmt.Sprintf("%s[%d]%s ", errMsg, idx+1, msg)
		}
	}
	panic(&GeneralError{
		HttpCode: httpCode,
		ErrCode:  8,
		ErrMsg:   errMsg,
	})
}

func PermissionDenied(httpCode uint32, errMsgs ...string) {
	errMsg := "permission denied:"
	if len(errMsgs) > 0 {
		errMsg = errMsg + ":"
		for idx, msg := range errMsgs {
			errMsg = fmt.Sprintf("%s[%d]%s ", errMsg, idx+1, msg)
		}
	}
	panic(&GeneralError{
		HttpCode: httpCode,
		ErrCode:  11,
		ErrMsg:   errMsg,
	})
}

func DuplicateOperation(httpCode uint32, errMsgs ...string) {
	errMsg := "duplicate denied:"
	if len(errMsgs) > 0 {
		errMsg = errMsg + ":"
		for idx, msg := range errMsgs {
			errMsg = fmt.Sprintf("%s[%d]%s ", errMsg, idx+1, msg)
		}
	}
	panic(&GeneralError{
		HttpCode: httpCode,
		ErrCode:  16,
		ErrMsg:   errMsg,
	})
}
