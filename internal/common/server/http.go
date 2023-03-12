package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

func RunHTTPServer(addr string, router *iris.Application) {
	router.Use(iris.Compression)
	router.Validator = validator.New()

	err := router.Listen(addr)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}
