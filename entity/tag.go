package entity

import "time"

type Tag struct {
	ID        int64
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
