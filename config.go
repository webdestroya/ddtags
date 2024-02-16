package ddtags

import "time"

type Config struct {
	// Datadog's MetricWithTimestamp function
	// it's probably ddlambda.MetricWithTimestamp
	MetricWithTimestamp func(string, float64, time.Time, ...string)

	// The string format passed to fmt.Sprintf for integer values
	// Default: '%d'
	IntegerFormat string

	// The string format passed to fmt.Sprintf for float values
	// Default: '%0.5f'
	FloatFormat string
}

var defaultConfig *Config

func Configure(cfg *Config) {
	defaultConfig = cfg
}

func init() {
	defaultConfig = &Config{
		FloatFormat:   "%.5f",
		IntegerFormat: "%d",
	}
}
