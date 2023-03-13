package router

import "github.com/kataras/iris/v12"

type HTTPHandler interface {
	Register(ctx iris.Context)
}

func NewHTTPRouter(handler HTTPHandler) *iris.Application {
	router := iris.New()

	userAPI := router.Party("/api/v1/user")
	{
		userAPI.Get("/register", handler.Register)
		userAPI.Put("/register", handler.Register)
	}

	return router
}
