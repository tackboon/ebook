package main

import (
	"github.com/tackboon/ebook/internal/common/server"
	"github.com/tackboon/ebook/internal/user/handler"
	"github.com/tackboon/ebook/internal/user/router"
	"github.com/tackboon/ebook/internal/user/service"
)

func main() {
	userService := service.NewUserService()
	httpServer := handler.NewHttpServer(userService)
	router := router.NewHTTPRouter(httpServer)
	server.RunHTTPServer(":5000", router)
}
