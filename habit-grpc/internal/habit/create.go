package habit

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Create creates and habit saves it and returns it
func Create(_ context.Context, h Habit) (*Habit, error) {
	h, err := ValidateAndFill(h)
	if err != nil {
		return nil, err
	}

	// todo: stored to datastore
	return &h, nil
}

// ValidateAndFill validates the habit and populates it with default values
// Returns Habit, InvalidInputError
func ValidateAndFill(h Habit) (Habit, error) {
	h.Name = Name(strings.TrimSpace(string(h.Name)))
	if h.Name == "" {
		return Habit{}, InvalidInputError{
			field:  "name",
			reason: "cannont be empty",
		}
	}

	if h.WeekelyFrequency == 0 {
		h.WeekelyFrequency = 1
	}

	if h.ID == "" {
		h.ID = ID(uuid.NewString())
	}

	if h.CreationTime.Equal(time.Time{}) {
		h.CreationTime = time.Now()
	}

	return h, nil
}
