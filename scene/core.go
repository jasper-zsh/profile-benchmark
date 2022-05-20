package scene

import "context"

type Scene interface {
	Run(ctx context.Context)
}
