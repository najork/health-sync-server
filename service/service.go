package service

import (
	"context"
	"sync"

	connect "github.com/abrander/garmin-connect"
	"github.com/najork/health-sync-server/conjure/api"
	"github.com/najork/health-sync-server/export"
	werror "github.com/palantir/witchcraft-go-error"
	"github.com/palantir/witchcraft-go-health/conjure/witchcraft/api/health"
	"github.com/palantir/witchcraft-go-logging/wlog/svclog/svc1log"
	"go.uber.org/atomic"
)

// compile-time assertion that Service implements api.HealthSyncService
var _ api.HealthSyncService = (*Service)(nil)

const (
	filename = "metrics.out"
)

type Service struct {
	client *connect.Client

	sync.RWMutex // guards health
	health       health.HealthCheckResult
	readiness    *atomic.Bool
}

func New(client *connect.Client) *Service {
	return &Service{
		client: client,
		health: health.HealthCheckResult{
			Type:  checkType,
			State: health.New_HealthState(health.HealthState_HEALTHY),
		},
		readiness: atomic.NewBool(false),
	}
}

func (s *Service) Start(ctx context.Context) error {
	if err := s.client.Authenticate(); err != nil {
		return err
	}
	s.readiness.Store(true)

	svc1log.FromContext(ctx).Info("Service started.")
	return nil
}

// Collect metrics from Garmin Connect for the given activity.
func (s *Service) Collect(ctx context.Context, requestArg api.ActivityRequest) error {
	if !s.readiness.Load() {
		return werror.ErrorWithContextParams(ctx, "service is not ready yet")
	}

	activity, err := s.client.Activity(requestArg.Id)
	if err != nil {
		return err
	}
	svc1log.FromContext(ctx).Debug("Retrieved activity.", svc1log.SafeParam("activity", activity))

	activityDetails, err := s.client.ActivityDetails(requestArg.Id, requestArg.MaxPoints)
	if err != nil {
		return err
	}
	svc1log.FromContext(ctx).Debug("Retrieved activity details.", svc1log.SafeParam("activityDetails", activityDetails))

	return export.ToTextFile(ctx, filename, activity, activityDetails)
}
