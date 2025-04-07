package internal

import "time"

type Post struct {
	ID           int
	AuthorID     int
	ReferralCode string
	Title        string
	Content      string
	CreatedAt    time.Time
}
