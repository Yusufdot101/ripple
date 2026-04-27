package grpc

import (
	userpb "github.com/Yusufdot101/ripple-proto/golang/user/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	userClient userpb.UserServiceClient
	conn       *grpc.ClientConn
}

func NewAdapter(url string) (*Adapter, error) {
	conn, err := grpc.NewClient(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := userpb.NewUserServiceClient(conn)
	return &Adapter{
		conn:       conn,
		userClient: client,
	}, nil
}

func (a *Adapter) Close() error {
	return a.conn.Close()
}
