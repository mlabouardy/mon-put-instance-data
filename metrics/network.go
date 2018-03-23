package metrics

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/cloudwatch/services"
	"github.com/shirou/gopsutil/net"
)

type Network struct{}

func (n Network) Collect(instanceId string, c CloudWatchService) {
	networkMetrics, err := net.IOCounters(false)
	if err != nil {
		log.Fatal(err)
	}

	for _, iocounter := range networkMetrics {
		dimensions := make([]cloudwatch.Dimension, 0)
		dimensionKey1 := "InstanceId"
		dimensions = append(dimensions, cloudwatch.Dimension{
			Name:  &dimensionKey1,
			Value: &instanceId,
		})
		dimensionKey2 := "IOCounter"
		dimensions = append(dimensions, cloudwatch.Dimension{
			Name:  &dimensionKey2,
			Value: &iocounter.Name,
		})

		bytesInData := constructMetricDatum("BytesIn", float64(iocounter.BytesRecv), cloudwatch.StandardUnitBytes, dimensions)
		c.Publish(bytesInData, "CustomMetrics")
		bytesOutData := constructMetricDatum("BytesOut", float64(iocounter.BytesSent), cloudwatch.StandardUnitBytes, dimensions)
		c.Publish(bytesOutData, "CustomMetrics")

		packetsInData := constructMetricDatum("PacketsIn", float64(iocounter.PacketsRecv), cloudwatch.StandardUnitBytes, dimensions)
		c.Publish(packetsInData, "CustomMetrics")
		packetsOutData := constructMetricDatum("PacketsOut", float64(iocounter.PacketsSent), cloudwatch.StandardUnitBytes, dimensions)
		c.Publish(packetsOutData, "CustomMetrics")

		errorsInData := constructMetricDatum("ErrorsIn", float64(iocounter.Errin), cloudwatch.StandardUnitBytes, dimensions)
		c.Publish(errorsInData, "CustomMetrics")
		errorsOutData := constructMetricDatum("ErrorsOut", float64(iocounter.Errout), cloudwatch.StandardUnitBytes, dimensions)
		c.Publish(errorsOutData, "CustomMetrics")

		log.Printf("Network - %s Bytes In/Out: %v/%v Packets In/Out: %v/%v Errors In/Out: %v/%v\n",
			iocounter.Name, iocounter.BytesRecv, iocounter.BytesSent, iocounter.Errin,
			iocounter.Errout, iocounter.PacketsRecv, iocounter.PacketsSent)
	}
}
