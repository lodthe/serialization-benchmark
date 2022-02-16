package bench

import (
	"fmt"
	"reflect"
	"time"

	"github.com/lodthe/serialization-benchmark/format"
	"github.com/lodthe/serialization-benchmark/sample"

	"github.com/pkg/errors"
)

type Runner struct {
	config Config
}

func NewRunner(config Config) *Runner {
	return &Runner{
		config: config,
	}
}

func (r *Runner) Bench(s format.Serializer) (res Result, err error) {
	err = r.benchMarshal(s, &res)
	if err != nil {
		return Result{}, errors.Wrap(err, "marshal benchmark failed")
	}

	err = r.validateUnmarshalling(s, &res)
	if err != nil {
		return Result{}, errors.Wrap(err, "unmarshal validation failed")
	}

	err = r.benchUnmarshal(s, &res)
	if err != nil {
		return Result{}, errors.Wrap(err, "unmarshal benchmark failed")
	}

	return res, nil
}

func (r *Runner) benchMarshal(s format.Serializer, res *Result) error {
	var totalDuration time.Duration
	benchStartedAt := time.Now()
	for time.Since(benchStartedAt) < r.config.DurationThreshold/2 {
		iterationStartedAt := time.Now()
		for i := 0; i < r.config.SerialRunCount; i++ {
			marshalled, err := s.Marshal(r.config.Data)
			if err != nil {
				return err
			}

			_ = marshalled
		}

		elapsed := time.Since(iterationStartedAt)
		totalDuration += elapsed

		res.MarshalExecCount += r.config.SerialRunCount
	}

	res.MarshalMeanDuration = time.Duration(int64(totalDuration) / int64(res.MarshalExecCount))

	return nil
}

func (r *Runner) validateUnmarshalling(s format.Serializer, res *Result) error {
	marshalled, err := s.Marshal(r.config.Data)
	if err != nil {
		return errors.Wrap(err, "Marshal failed")
	}

	res.MarshalledData = marshalled

	unmarshalled := new(sample.User)
	err = s.Unmarshal(res.MarshalledData, unmarshalled)
	if err != nil {
		return errors.Wrap(err, "Unmarshal failed")
	}

	unmarshalled.ToUTC()

	if !reflect.DeepEqual(r.config.Data, *unmarshalled) {
		fmt.Println(*unmarshalled)
		fmt.Println(r.config.Data)
		return errors.New("output differs from the input")
	}

	return nil
}

func (r *Runner) benchUnmarshal(s format.Serializer, res *Result) error {
	var totalDuration time.Duration
	benchStartedAt := time.Now()
	for time.Since(benchStartedAt) < r.config.DurationThreshold/2 {
		iterationStartedAt := time.Now()
		for i := 0; i < r.config.SerialRunCount; i++ {
			output := new(sample.User)
			err := s.Unmarshal(res.MarshalledData, output)
			if err != nil {
				return err
			}
		}

		elapsed := time.Since(iterationStartedAt)
		totalDuration += elapsed

		res.UnmarshalExecCount += r.config.SerialRunCount
	}

	res.UnmarshalMeanDuration = time.Duration(int64(totalDuration) / int64(res.UnmarshalExecCount))

	return nil
}
