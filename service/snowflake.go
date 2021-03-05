package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/w3xse7en/web-test/pkg/api"
)

type server struct {
	api.UnimplementedSnowflakeServer
}

func NewSnowflakeServer() *server {
	return &server{}
}

func (s *server) GetID(ctx context.Context, in *empty.Empty) (*api.GetIDResp, error) {
	return &api.GetIDResp{Id: 123}, nil
}

func (s *server) GetIDByRateLimit(ctx context.Context, in *empty.Empty) (*api.GetIDResp, error) {
	return &api.GetIDResp{Id: 123}, nil
}
