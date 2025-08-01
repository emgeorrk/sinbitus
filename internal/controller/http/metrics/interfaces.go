package metrics

import "time"

type TimeProvider interface {
	Now() time.Time
}
