package metrics

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/cloudwatch"
	"github.com/shirou/gopsutil/mem"
)

type Memory struct{}

func (m Memory) Collect(instanceId string, cloudWatch CloudWatch) {
	memoryMetrics, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	putMetricData(cfg, "CustomMetrics", "MemoryUtilization", memoryMetrics.UsedPercent, cloudwatch.StandardUnitPercent, "1-78787")
	putMetricData(cfg, "CustomMetrics", "MemoryUsed", float64(memoryMetrics.Used), cloudwatch.StandardUnitBytes, "1-78787")
	putMetricData(cfg, "CustomMetrics", "MemoryAvailable", float64(memoryMetrics.Available), cloudwatch.StandardUnitBytes, "1-78787")
}
