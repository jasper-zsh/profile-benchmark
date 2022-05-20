package main

import (
	"context"
	"fmt"
	"github.com/pyroscope-io/client/pyroscope"
	"profile-benchmark/scene"
	"time"
)

func main() {
	ctx := context.Background()
	s := scene.NewComputeScene(time.Second * 30)
	//s := scene.NewDiskIOScene(time.Second * 30)
	tps1, err := s.Run(ctx)
	if err != nil {
		panic(err)
	}

	pyroscope.Start(pyroscope.Config{
		ApplicationName: "benchmark",
		ServerAddress:   "http://localhost:4040",
		Logger:          pyroscope.StandardLogger,
		DisableGCRuns:   true,
	})

	tps2, err := s.Run(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Plain TPS: %.3f", tps1)
	fmt.Printf("Profiling TPS: %.3f\n", tps2)
	fmt.Printf("Diff: %.2f%%", (tps2-tps1)/tps1*100)
}
