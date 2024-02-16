package ddtags_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/webdestroya/ddtags"
)

func TestMetric(t *testing.T) {
	tags := &simpleTagStruct{
		TagA: "vala",
		TagB: "valb",
	}

	funcCalled := false
	expTagList := ddtags.Extract(tags)

	ddtags.Configure(&ddtags.Config{
		MetricWithTimestamp: func(mn string, v float64, ts time.Time, tagList ...string) {
			require.Equal(t, "metric_name", mn)
			require.Equal(t, 123.4567, v)
			require.InDelta(t, time.Now().Unix(), ts.Unix(), 1)
			require.ElementsMatch(t, expTagList, tagList)
			funcCalled = true
		},
	})

	ddtags.Metric("metric_name", 123.4567, tags)
	require.True(t, funcCalled)
}

func TestMetricWithTimestamp(t *testing.T) {
	tags := &simpleTagStruct{
		TagA: "vala",
		TagB: "valb",
	}

	funcCalled := false
	expTagList := ddtags.Extract(tags)

	timeNow := time.Now()

	ddtags.Configure(&ddtags.Config{
		MetricWithTimestamp: func(mn string, v float64, ts time.Time, tagList ...string) {
			require.Equal(t, "metric_name", mn)
			require.Equal(t, 123.4567, v)
			require.Equal(t, timeNow, ts)
			require.ElementsMatch(t, expTagList, tagList)
			funcCalled = true
		},
	})

	ddtags.MetricWithTimestamp("metric_name", 123.4567, timeNow, tags)
	require.True(t, funcCalled)

}
