package repository

import (
	"context"
	"habit-grpc/internal/habit"
)

type Error string

type Logger interface {
	Logf(format string, args ...any)
}

// HabitRepository is a database to store habits
type HabitRepository struct {
	lgr   Logger
	store map[habit.ID]habit.Habit
}

// New returns setups and a new HabitRepository
func New(lgr Logger) *HabitRepository {
	return &HabitRepository{
		lgr:   lgr,
		store: make(map[habit.ID]habit.Habit),
	}
}

func (hr *HabitRepository) Add(_ context.Context, habit habit.Habit) error {
	hr.lgr.Logf("storing new habit in repository: %v", habit)
	hr.store[habit.ID] = habit
	return nil
}

func (hr *HabitRepository) ListAll(_ context.Context) ([]habit.Habit, error) {
	hr.lgr.Logf("listing habits in repository")
	var res []habit.Habit
	for _, habit := range hr.store {
		res = append(res, habit)
	}
	return res, nil
}
