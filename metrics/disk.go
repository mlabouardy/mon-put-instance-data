package metrics

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/shirou/gopsutil/disk"
)

type Disk struct{}

func (d Disk) Collect(instanceId string) {
	diskMetrics, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}

	putMetricData(cfg, "CustomMetrics", "DiskUtilization", diskMetrics.UsedPercent, cloudwatch.StandardUnitPercent, "1-78787")
	putMetricData(cfg, "CustomMetrics", "DiskUsed", float64(diskMetrics.Used), cloudwatch.StandardUnitBytes, "1-78787")
	putMetricData(cfg, "CustomMetrics", "DiskFree", float64(diskMetrics.Free), cloudwatch.StandardUnitBytes, "1-78787")
}
