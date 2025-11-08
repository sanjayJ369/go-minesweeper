package habit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_validateAndFill(t *testing.T) {
	t.Run("Full", testValidateAndFillDetailsFull)
	t.Run("partial", testValidateAndFillPartial)
	t.Run("partial", testValidateAndFillSapceName)
}

func testValidateAndFillSapceName(t *testing.T) {
	t.Parallel()
	h := Habit{
		Name: "     ",
	}
	_, err := ValidateAndFill(h)
	assert.Error(t, err)
}

func testValidateAndFillPartial(t *testing.T) {
	t.Parallel()
	h := Habit{
		WeekelyFrequency: 1,
		Name:             "name",
	}

	got, err := ValidateAndFill(h)
	require.NoError(t, err)
	assert.NotEmpty(t, got.ID)
	assert.NotEmpty(t, got.CreationTime)
	assert.Equal(t, h.Name, got.Name)
	assert.Equal(t, h.WeekelyFrequency, got.WeekelyFrequency)
}

func testValidateAndFillDetailsFull(t *testing.T) {
	t.Parallel()
	h := Habit{
		ID:               "id",
		Name:             "name",
		CreationTime:     time.Now(),
		WeekelyFrequency: 1,
	}
	got, err := ValidateAndFill(h)
	require.NoError(t, err)
	assert.Equal(t, h, got)
}
