package service

import (
	"context"
	"net/http"

	"github.com/palantir/witchcraft-go-health/conjure/witchcraft/api/health"
)

const checkType = "service"

func (s *Service) HealthStatus(_ context.Context) health.HealthStatus {
	s.RLock()
	defer s.RUnlock()

	return health.HealthStatus{
		Checks: map[health.CheckType]health.HealthCheckResult{
			s.health.Type: s.health,
		},
	}
}

func (s *Service) Status() (respStatus int, metadata interface{}) {
	if !s.readiness.Load() {
		return http.StatusServiceUnavailable, nil
	}
	return http.StatusOK, nil
}
