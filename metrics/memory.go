package metrics

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/cloudwatch/services"
	"github.com/shirou/gopsutil/mem"
)

type Memory struct{}

func (m Memory) Collect(instanceId string, c CloudWatchService) {
	memoryMetrics, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	memoryUtilizationData := constructMetricDatum("MemoryUtilization", memoryMetrics.UsedPercent, cloudwatch.StandardUnitPercent, instanceId)
	c.Publish(memoryUtilizationData, "CustomMetrics")

	memoryUsedData := constructMetricDatum("MemoryUsed", float64(memoryMetrics.Used), cloudwatch.StandardUnitBytes, instanceId)
	c.Publish(memoryUsedData, "CustomMetrics")

	memoryAvailableData := constructMetricDatum("MemoryAvailable", float64(memoryMetrics.Available), cloudwatch.StandardUnitBytes, instanceId)
	c.Publish(memoryAvailableData, "CustomMetrics")

	log.Printf("Memory - Utilization:%v%% Used:%v Available:%v\n", memoryMetrics.UsedPercent, memoryMetrics.Used, memoryMetrics.Available)
}
