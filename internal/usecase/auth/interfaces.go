package auth

import "time"

type (
	TimeProvider interface {
		Now() time.Time
	}
)
