package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	NotAuthorized = errors.New("not authorized")
)

type ErrMsg struct {
	Status int         `json:"-"`
	Code   int         `json:"code,omitempty"`
	Text   string      `json:"text"`
	Detail interface{} `json:"detail,omitempty"`
}

func (e *ErrMsg) Error() string { return e.Text }

var (
	ErrBadInput = &ErrMsg{
		Status: http.StatusNotAcceptable,
		Text:   "无效的表单输入",
	}
	ErrBadSignup = &ErrMsg{
		Status: http.StatusNotAcceptable,
		Text:   "无效的登录表单",
	}
	ErrUsernameExist = &ErrMsg{
		Status: http.StatusNotAcceptable,
		Text:   "用户名已存在",
	}
	ErrEmailExist = &ErrMsg{
		Status: http.StatusNotAcceptable,
		Text:   "电子邮箱已被使用",
	}

	ErrBadLogin = &ErrMsg{
		Status: http.StatusNotAcceptable,
		Text:   "无效的用户名或密码",
	}
	ErrUserNotExist = &ErrMsg{
		Status: http.StatusNotAcceptable,
		Text:   "用户名不存在",
	}
	ErrWrongLogin = &ErrMsg{
		Status: http.StatusUnauthorized,
		Text:   "用户名或密码错误",
	}
	ErrAuth = &ErrMsg{
		Status: http.StatusUnauthorized,
		Text:   "认证失败",
	}
)

func ErrJSON(c *gin.Context, e *ErrMsg, err ...error) error {
	c.JSON(e.Status, e)

	if len(err) == 0 {
		return e
	}
	return err[0]
}

func InternalError(c *gin.Context, err error) {
	c.AbortWithError(http.StatusInternalServerError, err)
}
