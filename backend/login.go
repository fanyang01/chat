package main

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/fanyang01/chat/backend/jwt"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" binding:"required,max=32"`
	Password string `form:"password" binding:"required,max=64,alphanum"`
}

func login(c *gin.Context) {
	var f LoginForm
	if err := c.Bind(&f); err != nil {
		ErrJSON(c, ErrBadLogin, err)
		return
	}

	var hash sql.NullString
	if err := DB.SQL(
		`SELECT passhash FROM users WHERE username = $1`, f.Username,
	).QueryScalar(&hash); err != nil {
		InternalError(c, err)
		return
	} else if !hash.Valid {
		// ErrJSON(c, ErrUserNotExist)
		ErrJSON(c, ErrWrongLogin)
		return
	} else if err := bcrypt.CompareHashAndPassword(
		[]byte(hash.String), []byte(f.Password),
	); err != nil {
		ErrJSON(c, ErrWrongLogin, err)
		return
	}

	claims := jwt.Claims{UserClaimKey: f.Username}
	c.JSON(http.StatusOK, gin.H{
		"token": claims.Sign(),
	})
}
