package model

import "time"

// Article - Article model
type Article struct {
	ID         int
	Title      string
	Body       string
	Archived   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt time.Time
}
