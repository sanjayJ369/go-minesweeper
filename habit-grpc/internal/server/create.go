package server

import (
	"context"
	"errors"
	"habit-grpc/api"
	"habit-grpc/internal/habit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateHabit(ctx context.Context, req *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	s.lgr.Logf("create habit request received: %v", req)
	var freq uint
	if req.WeekelyFrequency != nil {
		freq = uint(*req.WeekelyFrequency)
	}

	h := habit.Habit{
		Name:             habit.Name(req.Name),
		WeekelyFrequency: habit.WeekelyFrequency(freq),
	}

	createdHabit, err := habit.Create(ctx, h)
	if err != nil {
		var invalidErr habit.InvalidInputError
		if errors.As(err, &invalidErr) {
			return nil, status.Error(codes.InvalidArgument, invalidErr.Error())
		}

		return nil, status.Errorf(codes.Internal, "could not save the habit %v: %s", h, err.Error())
	}

	return &api.CreateHabitResponse{
		Habit: &api.Habit{
			Id:               string(createdHabit.ID),
			Name:             string(createdHabit.Name),
			WeekelyFrequency: int32(createdHabit.WeekelyFrequency),
		},
	}, nil
}
