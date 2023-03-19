package main

import (
	"github.com/tackboon/ebook/internal/common/client"
	"github.com/tackboon/ebook/internal/common/server"
	"github.com/tackboon/ebook/internal/user/adapter"
	"github.com/tackboon/ebook/internal/user/handler"
	"github.com/tackboon/ebook/internal/user/router"
	"github.com/tackboon/ebook/internal/user/service"
)

func main() {
	authClient, closeAuthClient, err := client.NewAuthClient()
	if err != nil {
		panic(err)
	}
	defer closeAuthClient()
	authGRPC := adapter.NewAuthGRPC(authClient)

	userService := service.NewUserService(authGRPC)
	httpServer := handler.NewHttpServer(userService)

	router := router.NewHTTPRouter(httpServer)
	server.RunHTTPServer(":5000", router)
}
