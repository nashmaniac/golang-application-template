package usecases

import (
	"context"

	"github.com/nashmaniac/golang-application-template/adapters"
)

type usecases struct {
	PeristenceStore adapters.PeristenceStore
}

func NewUsecases(
	ctx context.Context,
	p adapters.PeristenceStore,
) (adapters.Usecases, error) {
	return &usecases{
		PeristenceStore: p,
	}, nil
}
