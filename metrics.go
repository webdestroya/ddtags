package ddtags

import "time"

func Metric(metric string, value float64, tags any) {
	MetricWithTimestamp(metric, value, time.Now(), tags)
}

func MetricWithTimestamp(metric string, value float64, timestamp time.Time, tags any) {
	tagSlice := Extract(tags)

	defaultConfig.MetricWithTimestamp(metric, value, timestamp, tagSlice...)
}
