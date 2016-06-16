package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type SignupForm struct {
	Name     string `form:"name" binding:"required,max=32"`
	Gender   string `form:"gender" binding:"required,eq=male|eq=female"`
	Email    string `form:"email" binding:"required,email"`
	Username string `form:"username" binding:"required,max=32"`
	Password string `form:"password" binding:"required,max=64,alphanum"`
}

func signup(c *gin.Context) {
	var f SignupForm
	if err := c.Bind(&f); err != nil {
		ErrJSON(c, ErrBadSignup, err)
		return
	}

	do := func() (err error, internal bool) {
		tx, err := DB.Begin()
		if err != nil {
			return err, true
		}
		defer tx.AutoRollback()

		var exist struct {
			username, email bool
		}
		if err := tx.SQL(
			`SELECT
				EXISTS(SELECT 1 FROM users WHERE username = $1) AS username,
				EXISTS(SELECT 1 FROM users WHERE profile->>'email' = $2) AS email`,
			f.Username, f.Email,
		).QueryScalar(&exist.username, &exist.email); err != nil {
			return err, true
		} else if exist.username {
			return ErrJSON(c, ErrUsernameExist), false
		} else if exist.email {
			return ErrJSON(c, ErrEmailExist), false
		}

		hash, err := bcrypt.GenerateFromPassword(
			[]byte(f.Password), bcrypt.DefaultCost,
		)
		if err != nil {
			return err, true
		}

		profile, err := json.Marshal(&Profile{
			Username: f.Username,
			Name:     f.Name,
			Gender:   f.Gender,
			Email:    f.Email,
		})
		if err != nil {
			return err, true
		}

		if _, err := tx.SQL(
			`INSERT INTO users (username, passhash, profile) VALUES ($1, $2, $3)`,
			f.Username, hash, profile,
		).Exec(); err != nil {
			return err, true
		}

		return tx.Commit(), true
	}

	if err, internal := do(); err != nil {
		if internal {
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"username": f.Username,
	})
}
