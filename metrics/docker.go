package metrics

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/cloudwatch/services"
	"github.com/shirou/gopsutil/docker"
)

type Docker struct{}

func (d Docker) Collect(instanceId string, c CloudWatchService) {
	containers, err := docker.GetDockerStat()
	if err != nil {
		log.Fatal(err)
	}
	for _, container := range containers {
		containerMemory, err := docker.CgroupMemDocker(container.ContainerID)
		if err != nil {
			log.Fatal(err)
		}

		containerMemoryData := constructMetricDatum("ContainerMemory", float64(containerMemory.MemUsageInBytes), cloudwatch.StandardUnitBytes, instanceId)
		c.Publish(containerMemoryData, "CustomMetrics")

		containerCPU, err := docker.CgroupCPUDocker(container.ContainerID)
		if err != nil {
			log.Fatal(err)
		}

		containerCPUUserData := constructMetricDatum("ContainerCPUUser", float64(containerCPU.User), cloudwatch.StandardUnitSeconds, instanceId)
		c.Publish(containerCPUUserData, "CustomMetrics")

		containerCPUSystemData := constructMetricDatum("ContainerCPUSystem", float64(containerCPU.System), cloudwatch.StandardUnitSeconds, instanceId)
		c.Publish(containerCPUSystemData, "CustomMetrics")

		log.Printf("Docker - Container:%s Memory:%v User:%v System:%v\n", container.Name, containerMemory.MemMaxUsageInBytes, containerCPU.User, containerCPU.System)
	}
}
