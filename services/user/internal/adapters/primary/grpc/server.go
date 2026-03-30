package grpc

import "github.com/Yusufdot101/ripple-proto/golang/user"

type Adapter struct {
	user.UnimplementedUserServiceServer
}
