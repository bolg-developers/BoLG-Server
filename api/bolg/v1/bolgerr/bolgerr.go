package bolgerr

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	Internal        = status.New(codes.Internal, "サーバーエラー").Err()
	InvalidArgument = status.New(codes.InvalidArgument, "クライアントエラー").Err()
)
