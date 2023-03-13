package server

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"
	"github.com/sirupsen/logrus"
)

func RunHTTPServer(addr string, router *iris.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	router.Validator = validator.New()
	router.UseRouter(iris.Compression)
	router.UseRouter(iris.NoCache)
	router.UseRouter(crs)
	router.UseRouter(requestid.New())
	router.UseRouter(recover.New())
	router.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		router.ServeHTTP(w, r)
	})

	err := router.Listen(addr)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}
