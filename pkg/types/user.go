package types

import "time"

type UserPostListView struct {
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
	Citation string    `json:"citation"`
	Summary  string    `json:"summary"`
}
