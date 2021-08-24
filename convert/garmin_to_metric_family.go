package convert

import (
	"strings"
	"time"

	connect "github.com/abrander/garmin-connect"
	dto "github.com/prometheus/client_model/go"
)

func GarminActivityToMetricFamilies(activity *connect.Activity, details *connect.ActivityDetails) []*dto.MetricFamily {
	if activity == nil || details == nil {
		return nil
	}

	timestampIndex := -1
	metricFamilies := make(map[int]*dto.MetricFamily, len(details.MetricsDescriptors))
	for _, metricsDescriptor := range details.MetricsDescriptors {
		if metricsDescriptor.Key == "directTimestamp" {
			timestampIndex = metricsDescriptor.MetricsIndex
			continue
		}

		metricName := metricsDescriptor.Key
		metricType := keyToMetricType(metricsDescriptor.Key)
		metricFamilies[metricsDescriptor.MetricsIndex] = &dto.MetricFamily{
			Name: &metricName,
			Type: &metricType,
		}
	}

	for _, activityDetailMetric := range details.ActivityDetailMetrics {
		timestamp := metricToTimestampMs(activityDetailMetric.Metrics[timestampIndex])
		for i, metric := range activityDetailMetric.Metrics {
			if i == timestampIndex {
				continue
			}

			metricFamily := metricFamilies[i]
			switch *metricFamily.Type {
			case dto.MetricType_GAUGE:
				gauge := garminMetricToGauge(metric, &timestamp)
				metricFamily.Metric = append(metricFamilies[i].Metric, gauge)
			case dto.MetricType_COUNTER:
				counter := garminMetricToCounter(metric, &timestamp)
				metricFamily.Metric = append(metricFamilies[i].Metric, counter)
			default:
				continue
			}
		}
	}

	return mapToSlice(metricFamilies)
}

func keyToMetricType(key string) dto.MetricType {
	if strings.HasPrefix(key, "direct") {
		return dto.MetricType_GAUGE
	} else if strings.HasPrefix(key, "sum") {
		return dto.MetricType_COUNTER
	} else {
		return dto.MetricType_UNTYPED
	}
}

func metricToTimestampMs(metric connect.Metric) int64 {
	if metric == nil {
		return time.Unix(0, 0).UnixMilli()
	}
	return int64(*metric)
}

func garminMetricToGauge(value *float64, timestampMs *int64) *dto.Metric {
	return &dto.Metric{
		Gauge: &dto.Gauge{
			Value: value,
		},
		TimestampMs: timestampMs,
	}
}

func garminMetricToCounter(value *float64, timestampMs *int64) *dto.Metric {
	return &dto.Metric{
		Counter: &dto.Counter{
			Value: value,
		},
		TimestampMs: timestampMs,
	}
}

func mapToSlice(m map[int]*dto.MetricFamily) []*dto.MetricFamily {
	s := make([]*dto.MetricFamily, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}
