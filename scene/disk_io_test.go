package scene

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiskIOScene_Run(t *testing.T) {
	s := NewDiskIOScene()
	tps, err := s.Run(nil)
	assert.NoError(t, err)
	fmt.Printf("Disk IO TPS: %.3f\n", tps)
}
