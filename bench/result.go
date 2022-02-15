package bench

import (
	"time"
)

type Result struct {
	MarshalledData []byte

	MarshalExecCount    int
	MarshalMeanDuration time.Duration

	UnmarshalExecCount    int
	UnmarshalMeanDuration time.Duration
}
