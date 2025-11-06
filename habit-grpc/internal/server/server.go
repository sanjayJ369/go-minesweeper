package server

import (
	"fmt"
	"habit-grpc/api"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Logger interface {
	Logf(format string, args ...any)
}

type Server struct {
	api.UnimplementedHabitsServer
	lgr Logger
}

func New(lgr Logger) *Server {
	return &Server{
		lgr: lgr,
	}
}

func (s *Server) ListenAndServer(port int) error {
	const addr = "127.0.0.1"

	listener, err := net.Listen("tcp",
		net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
		return fmt.Errorf("unable to listen at tcp port %d: %w", port, err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s)

	s.lgr.Logf("starting the server on port: %d", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("unable to server grpc server on port: %d: %w", port, err)
	}

	return nil
}
