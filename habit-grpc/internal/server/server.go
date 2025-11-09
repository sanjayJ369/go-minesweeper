package server

import (
	"context"
	"fmt"
	"habit-grpc/api"
	"habit-grpc/internal/habit"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Logger interface {
	Logf(format string, args ...any)
}

type Repository interface {
	Add(context.Context, habit.Habit) error
	ListAll(context.Context) ([]habit.Habit, error)
}

type Server struct {
	api.UnimplementedHabitsServer
	repo Repository
	lgr  Logger
}

func New(repo Repository, lgr Logger) *Server {
	return &Server{
		lgr:  lgr,
		repo: repo,
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
