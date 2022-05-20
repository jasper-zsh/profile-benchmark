package scene

import (
	"fmt"
	"github.com/pyroscope-io/client/pyroscope"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeScene_Run(t *testing.T) {
	pyroscope.Start(pyroscope.Config{
		ApplicationName: "benchmark",
		ServerAddress:   "http://localhost:4040",
		Logger:          pyroscope.StandardLogger,
		DisableGCRuns:   true,
	})
	s := &ComputeScene{}
	tps, err := s.Run(nil)
	assert.NoError(t, err)
	fmt.Printf("TPS Compute: %.3f\n", tps)
}
