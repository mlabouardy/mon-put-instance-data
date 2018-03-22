package metrics

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/cloudwatch/services"
	"github.com/shirou/gopsutil/disk"
)

type Disk struct{}

func (d Disk) Collect(instanceId string, c CloudWatchService) {
	diskMetrics, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}

	diskUtilizationData := constructMetricDatum("DiskUtilization", diskMetrics.UsedPercent, cloudwatch.StandardUnitPercent, instanceId)
	c.Publish(diskUtilizationData, "CustomMetrics")

	diskUsedData := constructMetricDatum("DiskUsed", float64(diskMetrics.Used), cloudwatch.StandardUnitBytes, instanceId)
	c.Publish(diskUsedData, "CustomMetrics")

	diskFreeData := constructMetricDatum("DiskFree", float64(diskMetrics.Free), cloudwatch.StandardUnitBytes, instanceId)
	c.Publish(diskFreeData, "CustomMetrics")

	log.Printf("Disk - Utilization:%v%% Used:%v Free:%v\n", diskMetrics.UsedPercent, diskMetrics.Used, diskMetrics.Free)
}
