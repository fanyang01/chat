package main

import (
	"net/http"
	"strings"

	"github.com/fanyang01/chat/backend/jwt"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func login(c *gin.Context) {
	var f LoginForm
	if err := c.Bind(&f); err != nil {
		ErrJSON(c, ErrBadLogin, err)
		return
	}

	// DB query...

	claims := jwt.Claims{"user": f.Username}
	c.JSON(http.StatusOK, gin.H{
		"token": claims.Sign(),
	})
}

func auth(c *gin.Context) {
	claims, err := verify(c)
	if err != nil {
		ErrJSON(c, ErrAuth, err)
		return
	}
	c.Set("user", claims["user"])
	c.Next()
}

func verify(c *gin.Context) (jwt.Claims, error) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		return nil, NotAuthorized
	}
	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, NotAuthorized
	}
	return jwt.Verify(parts[1])
}
