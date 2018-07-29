package rpc

import (
	"context"

	pb "github.com/msyrus/hello-go/proto/hello"
)

// DefaultGreeting returns default greetings msg
func (s *Server) DefaultGreeting(_ context.Context, req *pb.ReqGreet) (*pb.Greet, error) {
	msg, err := s.svc.GreetDefault(req.GetName())
	if err != nil {
		return nil, err
	}
	resp := pb.Greet{
		Message: msg,
	}
	return &resp, nil
}
