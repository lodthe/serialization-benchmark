package bench

import (
	"time"

	"github.com/lodthe/serialization-benchmark/sample"
)

type Config struct {
	// How long one format will be tested.
	DurationThreshold time.Duration

	// How many times should one format be tested without time measurement (fetching the system time is expensive).
	SerialRunCount int

	// Data for marshalling/unmarshalling.
	Data sample.User
}
