package main

import (
	"strings"

	"github.com/fanyang01/chat/backend/jwt"
	"github.com/gin-gonic/gin"
)

const (
	UserClaimKey = "user"
	UserCtxKey   = "username"
)

func auth(c *gin.Context) {
	claims, err := verify(c)
	if err != nil {
		ErrJSON(c, ErrAuth, err)
		return
	}
	c.Set(UserCtxKey, claims[UserClaimKey])
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
