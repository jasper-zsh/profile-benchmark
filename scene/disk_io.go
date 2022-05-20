package scene

import (
	"context"
	"math/rand"
	"os"
	"time"
)

var _ Scene = (*DiskIOScene)(nil)

type DiskIOScene struct {
	duration time.Duration
}

func NewDiskIOScene(duration time.Duration) *DiskIOScene {
	return &DiskIOScene{
		duration: duration,
	}
}

func (s DiskIOScene) Run(ctx context.Context) (float64, error) {
	fileSize := int64(1024 * 1024 * 50)
	buf := make([][]byte, 0)
	for i := 0; i < 1024; i++ {
		b := make([]byte, 4096)
		_, _ = rand.Read(b)
		buf = append(buf, b)
	}
	filename := "tmp_diskio"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	err = f.Truncate(fileSize)
	if err != nil {
		return 0, err
	}
	written := 0
	beginTime := time.Now()
	var t time.Time
	for {
		b, err := f.WriteAt(buf[rand.Intn(len(buf))], rand.Int63n(fileSize-int64(len(buf[0]))))
		if err != nil {
			return 0, err
		}
		written += b
		t = time.Now()
		if t.After(beginTime.Add(s.duration)) {
			break
		}
	}
	return float64(written) / float64(t.UnixMilli()-beginTime.UnixMilli()) * 1000, nil
}
