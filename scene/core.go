package scene

import "context"

type Scene interface {
	Run(ctx context.Context) (tps float64, err error)
}
