package client

import (
	"github.com/pkg/errors"
	"github.com/tackboon/ebook/internal/common/genproto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient() (client auth.GreeterClient, close func() error, err error) {
	grpcAddr := "ebook-auth-grpc:6000"
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env MOBILE_GRPC_ADDR")
	}

	opts, err := grpcDialOpts(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return auth.NewGreeterClient(conn), conn.Close, nil
}

func grpcDialOpts(grpcAddr string) ([]grpc.DialOption, error) {
	// if noTLS, _ := strconv.ParseBool(os.Getenv("GRPC_NO_TLS")); noTLS {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}, nil
	// }

	// systemRoots, err := x509.SystemCertPool()
	// if err != nil {
	// 	return nil, errors.Wrap(err, "cannot load root CA cert")
	// }
	// creds := credentials.NewTLS(&tls.Config{
	// 	RootCAs:    systemRoots,
	// 	MinVersion: tls.VersionTLS12,
	// })

	// return []grpc.DialOption{
	// 	grpc.WithTransportCredentials(creds),
	// }, nil
}
