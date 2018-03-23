package metrics

import (
	"log"

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
		log.Printf("Network - %s Bytes In/Out: %v/%v Packets In/Out: %v/%s Errors In/Out: %v/%v\n",
			iocounter.Name, iocounter.BytesRecv, iocounter.BytesSent, iocounter.Errin,
			iocounter.Errout, iocounter.PacketsRecv, iocounter.PacketsSent)
	}
}
