package adapters

import "context"

type PeristenceStore interface {
	CloseDB(ctx context.Context)
}
