package connect

import "fmt"

type ActivityDetails struct {
	ActivityId            int                    `json:"activityId"`
	MeasurementCount      int                    `json:"measurementCount"`
	MetricsCount          int                    `json:"metricsCount"`
	MetricsDescriptors    []MetricsDescriptor    `json:"metricDescriptors"`
	ActivityDetailMetrics []ActivityDetailMetric `json:"activityDetailMetrics"`
	GeoPolylineDTO        GeoPolylineDTO         `json:"geoPolylineDTO"`
	HeartRateDTOs         []HeartRateDTO         `json:"heartRateDTOs"`
	DetailsAvailable      bool                   `json:"detailsAvailable"`
}

type MetricsDescriptor struct {
	MetricsIndex int    `json:"metricsIndex"`
	Key          string `json:"key"`
	Unit         Unit   `json:"unit"`
}

type Unit struct {
	ID     int     `json:"id"`
	Key    string  `json:"key"`
	Factor float64 `json:"factor"`
}

type ActivityDetailMetric struct {
	Metrics []Metric `json:"metrics"`
}

type Metric *float64

type GeoPolylineDTO struct {
	StartPoint Point   `json:"startPoint"`
	EndPoint   Point   `json:"endPoint"`
	MinLat     float64 `json:"minLat"`
	MaxLat     float64 `json:"maxLat"`
	MinLon     float64 `json:"minLon"`
	MaxLon     float64 `json:"maxLon"`
	Polyline   []Point `json:"polyline"`
}

type Point struct {
	Lat                       float64     `json:"lat"`
	Lon                       float64     `json:"lon"`
	Altitude                  interface{} `json:"altitude"`
	Time                      int64       `json:"time"`
	TimerStart                bool        `json:"timerStart"`
	TimerStop                 bool        `json:"timerStop"`
	DistanceFromPreviousPoint interface{} `json:"distanceFromPreviousPoint"`
	DistanceInMeters          interface{} `json:"distanceInMeters"`
	Speed                     float64     `json:"speed"`
	CumulativeAscent          interface{} `json:"cumulativeAscent"`
	CumulativeDescent         interface{} `json:"cumulativeDescent"`
	ExtendedCoordinate        bool        `json:"extendedCoordinate"`
	Valid                     bool        `json:"valid"`
}

type HeartRateDTO struct {
	// TODO
}

func (c *Client) ActivityDetails(activityID int, maxChartSize int) (*ActivityDetails, error) {
	URL := fmt.Sprintf("https://connect.garmin.com/modern/proxy/activity-service/activity/%d/details?maxChartSize=%d",
		activityID,
		maxChartSize,
	)

	details := new(ActivityDetails)

	err := c.getJSON(URL, details)
	if err != nil {
		return nil, err
	}

	return details, nil
}
