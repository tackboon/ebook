package main

import (
	"context"
	"fmt"
	"net"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"github.com/tackboon/ebook/internal/auth/handler"
	"github.com/tackboon/ebook/internal/common/genproto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func main() {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			serverInterceptor,
		),
		grpc.ChainStreamInterceptor(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry),
		),
	)

	svc := handler.NewGRPCServer()
	auth.RegisterGreeterServer(grpcServer, svc)

	addr := ":6000"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("grpcEndpoint", addr).Info("Starting: gRPC Listener")
	logrus.Fatal(grpcServer.Serve(listen))

	// server := grpc.NewServer()
	// svc := handler.NewGRPCServer()
	// auth.RegisterGreeterServer(grpcServer, svc)

	// app := iris.New()
	// rootApp := mvc.New(app)
	// rootApp.Handle(&svc, mvc.GRPC{
	// 	Server:      server,               // Required.
	// 	ServiceName: "helloworld.Greeter", // Required.
	// 	Strict:      false,
	// })

	// app.Listen(":6000")
	// listen, err := net.Listen("tcp", ":6000")
	// if err != nil {
	// 	logrus.Panic(err)
	// }

	// logrus.WithField("grpcEndpoint", ":6000").Info("Starting GRPC listener")
	// err = server.Serve(listen)
	// if err != nil {
	// 	logrus.Panic(err)
	// }
}

func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Skip authorize when GetJWT is requested
	fmt.Println("hello interceptor")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("==")
	}
	fmt.Println(md)
	authHeader, ok := md["authorization"]
	if !ok {
		fmt.Println("___")
	}
	fmt.Println(authHeader)

	p, ok := peer.FromContext(ctx)
	if !ok {
		fmt.Println("777")
	}
	fmt.Println(p)

	// Calls the handler
	h, err := handler(ctx, req)

	return h, err
}
