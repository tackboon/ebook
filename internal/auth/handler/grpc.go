package handler

import (
	"context"

	"github.com/tackboon/ebook/internal/common/genproto/auth"
)

type GRPCServer struct {
	auth.UnimplementedGreeterServer
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

// func (g GRPCServer) UpdateDBVersion(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
// 	err := g.mobileService.UpdateDBVersion(ctx)
// 	return &empty.Empty{}, err
// }

// func (g GRPCServer) UpdateFileVersion(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
// 	err := g.mobileService.UpdateFileVersion(ctx)
// 	return &empty.Empty{}, err
// }

func (g GRPCServer) SayHello(ctx context.Context, r *auth.HelloRequest) (*auth.HelloReply, error) {
	return &auth.HelloReply{Message: "Hello " + r.GetName()}, nil
}
