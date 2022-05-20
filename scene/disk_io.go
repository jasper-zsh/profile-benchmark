package scene

import (
	"context"
)

var _ Scene = (*DiskIOScene)(nil)

type DiskIOScene struct{}

func NewDiskIOScene() *DiskIOScene {
	return &DiskIOScene{}
}

func (d DiskIOScene) Run(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
