// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package mysqldata

import (
	"time"
)

type Task struct {
	ID          int32     `json:"id"`
	Description string    `json:"description"`
	Assignee    string    `json:"assignee"`
	IsDone      bool      `json:"is_done"`
	DeadlineAt  time.Time `json:"deadline_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
