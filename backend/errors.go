package main

import (
	"errors"
	"log"
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
	ErrBadLogin = &ErrMsg{
		Status: http.StatusNotAcceptable,
		Text:   "无效的用户名或密码",
	}
	ErrAuth = &ErrMsg{
		Status: http.StatusUnauthorized,
		Text:   "认证失败",
	}
)

func ErrJSON(c *gin.Context, e *ErrMsg, err ...error) {
	if len(err) == 1 {
		log.Println(err)
	}
	c.JSON(e.Status, e)
}
