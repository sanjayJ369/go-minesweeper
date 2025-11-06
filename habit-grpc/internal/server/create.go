package server

import (
	"context"
	"habit-grpc/api"
)

func (s *Server) CreateHabit(_ context.Context, req *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	s.lgr.Logf("create habit request received: %v", req)
	return &api.CreateHabitResponse{
		Habit: &api.Habit{},
	}, nil
}
