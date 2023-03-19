package router

import (
	"github.com/kataras/iris/v12"
	"github.com/tackboon/ebook/internal/user/handler"
)

func NewHTTPRouter(handler handler.HTTPServer) *iris.Application {
	router := iris.New()

	userAPI := router.Party("/api/v1/user")
	{
		userAPI.Get("/profile", handler.GetProfile)
	}

	return router
}
