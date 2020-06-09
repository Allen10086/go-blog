package models

import "time"

type Article struct {
	Title       string    `db:"title"`
	Status      bool      `db:"status"`
	Comments    int64     `db:"comments"`
	View        int64     `db:"view"`
	Cover       string    `db:"cover"`
	Create_time time.Time `db:"create_time"`
	Update_time time.Time `db:"update_time"`
	Summary     string    `db:"summary"`
}
