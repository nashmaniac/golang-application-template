package usecases

import (
	"context"

	"github.com/nashmaniac/golang-application-template/adapters"
	"github.com/nashmaniac/golang-application-template/config"
)

type usecases struct {
	PeristenceStore adapters.PeristenceStore
	Configuration   *config.Config
}

func NewUsecases(
	ctx context.Context,
	p adapters.PeristenceStore,
	configuration *config.Config,
) (adapters.Usecases, error) {
	return &usecases{
		Configuration:   configuration,
		PeristenceStore: p,
	}, nil
}
