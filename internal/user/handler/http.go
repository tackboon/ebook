package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/tackboon/ebook/internal/user/models"
	"github.com/tackboon/ebook/internal/user/service"
)

type HTTPServer struct {
	userService service.UserService
}

func NewHttpServer(userService service.UserService) HTTPServer {
	return HTTPServer{
		userService: userService,
	}
}

func (h HTTPServer) GetProfile(ctx iris.Context) {
	var params models.GetProfileRequest
	err := ctx.ReadQuery(&params)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			fmt.Println(errs.Error())
		}
		ctx.Write([]byte("params_error"))
		return
	}
	h.userService.SayHello(ctx.Request().Context())
	ctx.Write([]byte("registers"))
}
