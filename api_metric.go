package metrics_lib

type ApiMetricLabels struct {
	Service string
	Vendor  string
	Call    string
	Result  Result
	Env     string
}

func ApiMetric(client Client, value float64, labels ApiMetricLabels) error {
	err := client.Histogram("api_request_duration_histogram_milliseconds", value, map[string]string{
		"service": labels.Service,
		"vendor":  labels.Vendor,
		"call":    labels.Call,
		"result":  labels.Result,
		"env":     labels.Env,
	}, 1)

	return err
}
