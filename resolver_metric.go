package metrics_lib

type Result = string

type ResolverMetricLabels struct {
	Resolver string
	Service  string
	Protocol string
	Result   Result
	Env      string
}

func ResolverMetric(client Client, value float64, labels ResolverMetricLabels) error {
	err := client.Histogram("resolver_request_duration_histogram_milliseconds", value, map[string]string{
		"resolver": labels.Resolver,
		"service":  labels.Service,
		"protocol": labels.Protocol,
		"result":   labels.Result,
		"env":      labels.Env,
	}, 1)

	return err
}
