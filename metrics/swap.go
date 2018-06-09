package metrics

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/cloudwatch/services"
	"github.com/shirou/gopsutil/mem"
)

// Swap metric entity
type Swap struct{}

// Collect Swap usage
func (d Swap) Collect(instanceID string, c CloudWatchService, namespace string) {
	swapMetrics, err := mem.SwapMemory()
	if err != nil {
		log.Fatal(err)
	}

	dimensionKey := "InstanceId"
	dimensions := []cloudwatch.Dimension{
		cloudwatch.Dimension{
			Name:  &dimensionKey,
			Value: &instanceID,
		},
	}

	swapUtilizationData := constructMetricDatum("SwapUtilization", swapMetrics.UsedPercent, cloudwatch.StandardUnitPercent, dimensions)
	c.Publish(swapUtilizationData, namespace)

	swapUsedData := constructMetricDatum("SwapUsed", float64(swapMetrics.Used), cloudwatch.StandardUnitBytes, dimensions)
	c.Publish(swapUsedData, namespace)

	swapFreeData := constructMetricDatum("SwapFree", float64(swapMetrics.Free), cloudwatch.StandardUnitBytes, dimensions)
	c.Publish(swapFreeData, namespace)

	log.Printf("Swap - Utilization:%v%% Used:%v Free:%v\n", swapMetrics.UsedPercent, swapMetrics.Used, swapMetrics.Free)
}
