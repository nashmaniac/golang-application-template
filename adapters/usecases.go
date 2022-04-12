package adapters

import "context"

type Usecases interface {
	GetHealthz(ctx context.Context, version string) (map[string]string, error)
}
