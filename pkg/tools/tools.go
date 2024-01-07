package main

import (
	"database/sql"
	"time"

	"gopkg.in/guregu/null.v3"
)

func NullTimeToTime(nt null.Time) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

func SqlNullTimeToTime(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
