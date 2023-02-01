/**
 * @author:       wangxuebing
 * @fileName:     errorx.go
 * @date:         2023/1/14 17:10
 * @description:
 */

package errorx

import (
	"errors"
	"gorm.io/gorm"
)

const (
	defaultCode             = 1001 //默认错误
	databaseErrCode         = 1002 //数据库查询错误
	databaseErrNoneDataCode = 1003 //数据库查询数据为空错误
)

type CodeErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeErr{
		Code: code,
		Msg:  msg,
	}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func NewDatabaseError() error {
	return NewCodeError(databaseErrCode, "操作失败[Database Error]")
}

func NewDatabaseNotFoundError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NewCodeError(databaseErrNoneDataCode, "查询[Database]数据为Null")
	}
	return NewCodeError(databaseErrCode, "操作失败[Database Error]")
}

func (e *CodeErr) Error() string {
	return e.Msg
}

func (e *CodeErr) Data() *CodeErrResp {
	return &CodeErrResp{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
