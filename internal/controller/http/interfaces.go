package http

import "time"

type TimeProvider interface {
	Now() time.Time
}
