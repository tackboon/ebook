package adapter

import (
	"context"
	"fmt"

	"github.com/tackboon/ebook/internal/common/genproto/auth"
)

type AuthGRPC struct {
	client auth.GreeterClient
}

func NewAuthGRPC(client auth.GreeterClient) AuthGRPC {
	return AuthGRPC{
		client: client,
	}
}

func (m AuthGRPC) SayHello(ctx context.Context) error {
	params := auth.HelloRequest{Name: "ball"}
	r, err := m.client.SayHello(ctx, &params)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
