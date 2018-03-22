package metrics

import (
	. "github.com/mlabouardy/cloudWatch-custom-metrics"
)

type Metric interface {
	Collect(string, CloudWatch)
}
