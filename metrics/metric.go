package metrics

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/mlabouardy/cloudwatch/services"
)

// Metric entity
type Metric interface {
	Collect(string, CloudWatchService)
}

// constructMetricDatum construct cloudwatch data object
func constructMetricDatum(metricName string, value float64, unit cloudwatch.StandardUnit, dimensions []cloudwatch.Dimension) []cloudwatch.MetricDatum {
	return []cloudwatch.MetricDatum{
		cloudwatch.MetricDatum{
			MetricName: &metricName,
			Dimensions: dimensions,
			Unit:       unit,
			Value:      &value,
		},
	}
}
