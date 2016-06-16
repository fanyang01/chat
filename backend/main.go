package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/mgutz/dat.v1"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

var (
	DB *runner.DB
)

func init() {
	db, err := sql.Open("postgres", "dbname=chat user=root password=secretTERCES host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	runner.MustPing(db)

	dat.EnableInterpolation = true
	runner.LogQueriesThreshold = 10 * time.Millisecond

	DB = runner.NewDB(db, "postgres")
}

func main() {
	g := gin.Default()

	g.POST("/api/login/", login)
	g.POST("/api/signup/", signup)

	_ = g.Group("/api/", auth)

	g.Run(":8799")
}
