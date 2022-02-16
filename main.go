package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/lodthe/serialization-benchmark/bench"
	"github.com/lodthe/serialization-benchmark/format"
	"github.com/lodthe/serialization-benchmark/format/fgob"
	"github.com/lodthe/serialization-benchmark/format/fjson"
	"github.com/lodthe/serialization-benchmark/format/fyaml"
	"github.com/lodthe/serialization-benchmark/sample"
)

func main() {
	// Allow to use only one CPU unit, because some implementations may use goroutines.
	runtime.GOMAXPROCS(1)

	runner := bench.NewRunner(bench.Config{
		DurationThreshold: 3 * time.Second,
		SerialRunCount:    50,
		Data:              sample.Sample,
	})

	runBench(runner, "JSON", fjson.NewSerializer())
	runBench(runner, "GOB", fgob.NewSerializer())
	runBench(runner, "YAML", fyaml.NewSerializer())
}

func runBench(runner *bench.Runner, formatName string, s format.Serializer) {
	res, err := runner.Bench(s)
	if err != nil {
		log.Fatalf("Bench for %s failed: %v\n", formatName, err)
	}

	formatDuration := func(d time.Duration) string {
		return fmt.Sprintf("%d µs", d.Microseconds())
	}

	fmt.Printf("%s:\n", formatName)
	fmt.Printf("\tMarshalled data size: %d\n", len(res.MarshalledData))
	fmt.Printf("\tMean Marshal duration: %s\n", formatDuration(res.MarshalMeanDuration))
	fmt.Printf("\tMean Unmarshal duration: %s\n\n\n", formatDuration(res.UnmarshalMeanDuration))
}
