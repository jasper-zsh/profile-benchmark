package scene

import (
	"context"
	"math"
	"math/rand"
	"time"
)

var _ Scene = (*ComputeScene)(nil)

type ComputeScene struct {
	duration time.Duration
}

func NewComputeScene(duration time.Duration) *ComputeScene {
	return &ComputeScene{
		duration: duration,
	}
}

func (s ComputeScene) Run(ctx context.Context) (tps float64, err error) {
	start := time.Now()

	//np := runtime.NumCPU()
	np := 1
	c := make(chan float64, np)

	con, cancel := context.WithCancel(ctx)
	for i := 0; i < np; i++ {
		go calc(con, c)
	}

	cnt := 0
	var r float64
	var end time.Time
	for {
		select {
		case <-con.Done():
			return float64(cnt) / float64(end.UnixMilli()-start.UnixMilli()), nil
		default:
			r += <-c
			cnt++
			end = time.Now()
			if end.After(start.Add(s.duration)) {
				cancel()
			}
		}
	}
}

func calc(ctx context.Context, ch chan float64) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- math.Sqrt(rand.Float64()):
		}
	}
}
