package service

import (
	"context"
	"sync"

	"github.com/najork/health-sync-server/conjure/api"
	"github.com/palantir/witchcraft-go-health/conjure/witchcraft/api/health"
)

// compile-time assertion that Service implements api.ApolloService
var _ api.ApolloService = (*Service)(nil)

type Service struct {
	sync.RWMutex // guards health
	health       health.HealthCheckResult
}

func New() *Service {
	return &Service{
		health: health.HealthCheckResult{
			Type:  checkType,
			State: health.New_HealthState(health.HealthState_HEALTHY),
		},
	}
}

// Collect triggers metrics collection from the given provider.
func (s *Service) Collect(ctx context.Context, provider api.Provider, request api.ProviderRequest) error {
	return nil
}
