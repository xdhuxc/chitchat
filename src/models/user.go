package models

import "time"

type Thread struct {
	ID        int
	UUID      string
	Topic     string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
