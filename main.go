package main

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/shirou/gopsutil/mem"
	"github.com/urfave/cli"
)

const DIMENSION_NAME = "InstanceId"

func GetInstanceId() {

}

func collectMetrics(cfg aws.Config, collectMemory, collectSwap, collectDisk bool) {
	swapMetrics, err := mem.SwapMemory()
	if err != nil {
		log.Fatal(err)
	}

	if collectMemory {

	}

	if collectSwap {
		putMetricData(cfg, "CustomMetrics", "SwapUtilization", swapMetrics.UsedPercent, cloudwatch.StandardUnitPercent, "1-78787")
		putMetricData(cfg, "CustomMetrics", "SwapUsed", float64(swapMetrics.Used), cloudwatch.StandardUnitBytes, "1-78787")
	}

	if collectDisk {

	}
}

func main() {
	app := cli.NewApp()
	app.Name = "CloudWatch"
	app.Usage = "Publish Custom Metrics to CloudWatch"
	app.Version = "1.0.0"
	app.Author = "Mohamed Labouardy"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "memory",
			Usage: "Collect memory metrics",
		},
		cli.BoolFlag{
			Name:  "swap",
			Usage: "Collect swap metrics",
		},
		cli.BoolFlag{
			Name:  "disk",
			Usage: "Collect disk metrics",
		},
		cli.IntFlag{
			Name:  "interval",
			Usage: "Time interval",
			Value: 5,
		},
	}
	app.Action = func(c *cli.Context) error {
		collectMemory := c.Bool("memory")
		collectSwap := c.Bool("swap")
		collectDisk := c.Bool("disk")
		interval := c.Int("interval")

		cfg, err := external.LoadDefaultAWSConfig()
		if err != nil {
			panic("Unable to load SDK config")
		}

		duration := time.Duration(interval) * time.Minute
		for _ = range time.Tick(duration) {
			collectMetrics(cfg, collectMemory, collectSwap, collectDisk)
			log.Printf("Waiting for %d minutes ...", interval)
		}

		return nil
	}
	app.Run(os.Args)
}
