package ddtags

import "time"

type Config struct {
	// Datadog's MetricWithTimestamp function
	// it's probably ddlambda.MetricWithTimestamp
	MetricWithTimestamp func(string, float64, time.Time, ...string)

	// Precision for FormatFloat
	// Default: -1
	FloatPrecision int

	// BitSize for FormatFloat
	// Default: 64
	FloatBitSize int
}

var defaultConfig *Config

func Configure(cfg *Config) {
	defaultConfig = cfg
}

func init() {
	defaultConfig = &Config{
		FloatPrecision: -1,
		FloatBitSize:   64,
	}
}
