package model

import "time"

type Gathering struct {
	ID          int64     `db:"id"`
	Creator     int64     `db:"creator"`
	Location    string    `db:"location"`
	ScheduledAt time.Time `db:"scheduled_at"`
	Name        string    `db:"name"`
	Type        string    `db:"type"`
}
