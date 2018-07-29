package rpc

import (
	"github.com/msyrus/hello-go/service"
)

// Server represets the promo code rpc server
type Server struct {
	svc *service.Greeting
}

// NewServer returns a new grpc server instance
func NewServer(svc *service.Greeting) *Server {
	return &Server{
		svc: svc,
	}
}

// var unaryInterceptors = []grpc.UnaryServerInterceptor{}

// srvr := grpc.NewServer(
// 	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
// )
// ptypes.RegisterGreetingsServer(srvr, &Server{})
// reflection.Register(srvr)
