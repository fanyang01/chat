package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	g.POST("/api/login/", login)

	g.Run(":8799")
}
