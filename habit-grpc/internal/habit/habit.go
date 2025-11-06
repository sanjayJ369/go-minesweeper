package habit

import "time"

type ID string

type Name string

type WeekelyFrequency uint

type Habit struct {
	ID               ID
	Name             Name
	WeekelyFrequency WeekelyFrequency
	CreationTime     time.Time
}
