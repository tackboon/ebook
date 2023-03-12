package handler

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/kataras/iris/v12"
	"github.com/tackboon/ebook/internal/user/router"
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

func (h HTTPServer) Register(ctx iris.Context) {
	var user router.UserRequest
	err := ctx.ReadJSON(&user)
	if err != nil {
		// Handle the error, below you will find the right way to do that...

		if errs, ok := err.(validator.ValidationErrors); ok {
			// Wrap the errors with JSON format, the underline library returns the errors as interface.
			// validationErrors := wrapValidationErrors(errs)
			fmt.Println(errs)

			// Fire an application/json+problem response and stop the handlers chain.
			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
				Title("Validation error").
				Detail("One or more fields failed to be validated").
				Type("/user/validation-errors").
				Key("errors", "alskdjf"))

			return
		}

		// It's probably an internal JSON error, let's dont give more info here.
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}
	ctx.Write([]byte("register"))
}
