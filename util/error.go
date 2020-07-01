package util

import (
	"errors"
	"github.com/shao1f/user.tag.logic/model"
	"strings"
)

var (
	ErrIsNil          = errors.New("操作成功")
	ErrParamError     = errors.New("参数错误")
	ErrInternalError  = errors.New("内部错误")
	ErrInsertTagError = errors.New("新增标签错误")
)

var errMap = map[error]int{
	ErrIsNil: 0,
}

func Result(err error, errMsg ...string) model.BaseResponse {
	if err == nil {
		err = ErrIsNil
	}
	_, exists := errMap[err]
	if !exists {
		err = ErrInternalError
	}
	code := errMap[err]
	msg := err.Error()
	if len(msg) > 0 {
		msg = strings.Join(errMsg, ",")
	}
	resp := model.BaseResponse{
		ErrCode: code,
		ErrMsg:  msg,
	}
	return resp
}
