package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RawJSON(c *gin.Context, v sql.NullString) {
	c.Data(http.StatusOK, "application/json", []byte(v.String))
}

type Profile struct {
	Username string `json:"username,omitempty" form:"username" binding:"max=32"`
	Name     string `json:"name,omitempty" form:"name" binding:"max=32"`
	Gender   string `json:"gender,omitempty" form:"gender" binding:"omitempty,eq=male|eq=female"`
	Email    string `json:"email,omitempty" form:"email" binding:"omitempty,email"`
}

func selfProfile(c *gin.Context) {
	v, _ := c.Get(UserCtxKey)
	username := v.(string)

	var p sql.NullString
	if err := DB.SQL(
		`SELECT profile FROM users WHERE username = $1`, username,
	).QueryScalar(&p); err != nil {
		InternalError(c, err)
		return
	} else if !p.Valid {
		ErrJSON(c, ErrUserNotExist)
		return
	}
	RawJSON(c, p)
}

func updateProfile(c *gin.Context) {
	v, _ := c.Get(UserCtxKey)
	username := v.(string)

	var prof Profile
	if err := c.Bind(&prof); err != nil {
		ErrJSON(c, ErrBadInput, err)
		return
	}

	b, err := json.Marshal(&prof)
	if err != nil {
		InternalError(c, err)
		return
	}

	var p sql.NullString
	if err := DB.SQL(
		`UPDATE users SET profile = profile || $1 WHERE username = $2 RETURNING profile`,
		b, username,
	).QueryScalar(&p); err != nil {
		InternalError(c, err)
		return
	} else if !p.Valid {
		ErrJSON(c, ErrUserNotExist)
		return
	}
	RawJSON(c, p)
}

func friendProfile(c *gin.Context) {
}
