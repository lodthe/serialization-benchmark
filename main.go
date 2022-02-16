package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/lodthe/serialization-benchmark/bench"
	"github.com/lodthe/serialization-benchmark/format"
	"github.com/lodthe/serialization-benchmark/format/fgob"
	"github.com/lodthe/serialization-benchmark/format/fjson"
	"github.com/lodthe/serialization-benchmark/format/fxml"
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

	runBench(runner, "xml", fxml.NewSerializer())
	runBench(runner, "json", fjson.NewSerializer())
	runBench(runner, "gdb", fgob.NewSerializer())
	runBench(runner, "yaml", fyaml.NewSerializer())
}

func runBench(runner *bench.Runner, formatName string, s format.Serializer) {
	res, err := runner.Bench(s)
	if err != nil {
		log.Fatalf("bench for %s failed: %v\n", formatName, err)
	}

	formatDuration := func(d time.Duration) string {
		return fmt.Sprintf("%d µs", d.Microseconds())
	}

	fmt.Printf("%s:\n", formatName)
	fmt.Printf("\tMarshalled data size: %d\n", len(res.MarshalledData))
	fmt.Printf("\tMean Marshal duration: %s\n", formatDuration(res.MarshalMeanDuration))
	fmt.Printf("\tMean Unmarshal duration: %s\n\n\n", formatDuration(res.UnmarshalMeanDuration))

	saveMarshalledData(res.MarshalledData, formatName)
}

func saveMarshalledData(data []byte, format string) {
	file, err := os.Create(path.Join("output", format+".txt"))
	if err != nil {
		log.Fatalf("failed to create output file for %s: %v\n", format, err)
	}
	defer file.Close()

	_, err = fmt.Fprint(file, string(data))
	if err != nil {
		log.Fatalf("failed to write output file for %s: %v\n", format, err)
	}
}
