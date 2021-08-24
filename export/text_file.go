package export

import (
	"context"
	"os"

	connect "github.com/abrander/garmin-connect"
	"github.com/najork/health-sync-server/convert"
	"github.com/palantir/witchcraft-go-logging/wlog/svclog/svc1log"
	"github.com/prometheus/common/expfmt"
)

const (
	perm = 0600
)

func ToTextFile(ctx context.Context, filename string, activity *connect.Activity, details *connect.ActivityDetails) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			svc1log.FromContext(ctx).Error("Failed to close file.", svc1log.Stacktrace(err))
		}
	}()

	metricFamilies := convert.GarminActivityToMetricFamilies(activity, details)
	for _, metricFamily := range metricFamilies {
		if _, err := expfmt.MetricFamilyToText(f, metricFamily); err != nil {
			svc1log.FromContext(ctx).Warn("Failed to export metric family.",
				svc1log.SafeParam("name", metricFamily.Name),
				svc1log.Stacktrace(err))
			continue
		}
	}
	if _, err := f.WriteString("# EOF"); err != nil {
		return err
	}
	return nil
}
