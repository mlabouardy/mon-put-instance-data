package metrics

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/mon-put-instance-data/services"
	"github.com/shirou/gopsutil/disk"
)

// Disk metric entity
type Disk struct{}

// Collect Disk used & free space
func (d Disk) Collect(instanceID string, c CloudWatchService, namespace string) {
	diskMetrics, err := disk.Usage("/")
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

	diskUtilizationData := constructMetricDatum("DiskUtilization", diskMetrics.UsedPercent, cloudwatch.StandardUnitPercent, dimensions)
	c.Publish(diskUtilizationData, namespace)

	diskUsedData := constructMetricDatum("DiskUsed", float64(diskMetrics.Used), cloudwatch.StandardUnitBytes, dimensions)
	c.Publish(diskUsedData, namespace)

	diskFreeData := constructMetricDatum("DiskFree", float64(diskMetrics.Free), cloudwatch.StandardUnitBytes, dimensions)
	c.Publish(diskFreeData, namespace)

	log.Printf("Disk - Utilization:%v%% Used:%v Free:%v\n", diskMetrics.UsedPercent, diskMetrics.Used, diskMetrics.Free)
}
