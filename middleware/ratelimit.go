package middleware

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type middleware struct {
}

func New() *middleware {
	return &middleware{}
}

func (m *middleware) RateLimit(filterUrls ...string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		for _, url := range filterUrls {
			if info.FullMethod != url {
				continue
			}
			p, ok := peer.FromContext(ctx)
			if !ok {
				continue
			}
			log.Printf("[RateLimit] match url:%v ip:%v", url, p.Addr.String())
		}
		return handler(ctx, req)
	}
}
