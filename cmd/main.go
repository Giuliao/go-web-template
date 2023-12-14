package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	"my.net/module/api"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(ginlogrus.Logger(logrus.New()))

	api.Register(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
