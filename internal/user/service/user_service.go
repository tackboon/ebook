package service

import (
	"context"
	"fmt"

	"github.com/tackboon/ebook/internal/user/adapter"
	"google.golang.org/grpc/metadata"
)

type UserService struct {
	authGRPC adapter.AuthGRPC
}

func NewUserService(client adapter.AuthGRPC) UserService {
	return UserService{authGRPC: client}
}

func (u UserService) SayHello(ctx context.Context) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "v4")

	err := u.authGRPC.SayHello(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
