package protobuf

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// server is used to implement TestServiceServer.
type server struct {
	okText string
}

func (s server) Echo(ctx context.Context, empty *emptypb.Empty) (*Ok, error) {
	return &Ok{Ok: s.okText}, nil
}

func (s server) Smth(ctx context.Context, empty *emptypb.Empty) (*Something, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) mustEmbedUnimplementedTestServiceServer() {
	// TODO implement me
	panic("implement me")
}

func StartServer(port int) error {
	lsn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("can't listen to port %d", port)
	}

	s := grpc.NewServer()
	RegisterTestServiceServer(s, server{okText: "ok"})
	if err := s.Serve(lsn); err != nil {
		return fmt.Errorf("can't listen to port %d", port)
	}
	return nil
}
