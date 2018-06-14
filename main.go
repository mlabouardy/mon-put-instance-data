package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	. "github.com/mlabouardy/mon-put-instance-data/metrics"
	. "github.com/mlabouardy/mon-put-instance-data/services"
	"github.com/urfave/cli"
)

// GetInstanceID return EC2 instance id
func GetInstanceID() (string, error) {
	value := os.Getenv("AWS_INSTANCE_ID")
	if len(value) > 0 {
		return value, nil
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://169.254.169.254/latest/meta-data/instance-id", nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Collect metrics about enabled metric
func Collect(metrics []Metric, c CloudWatchService) {
	id, err := GetInstanceID()
	if err != nil {
		log.Fatal(err)
	}
	for _, metric := range metrics {
		metric.Collect(id, c)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "mon-put-instance-data"
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
		cli.BoolFlag{
			Name:  "network",
			Usage: "Collect network metrics",
		},
		cli.BoolFlag{
			Name:  "docker",
			Usage: "Collect containers metrics",
		},
		cli.StringFlag{
			Name:  "region",
			Usage: "AWS region",
			Value: "us-east-1",
		},
		cli.IntFlag{
			Name:  "interval",
			Usage: "Time interval",
			Value: 5,
		},
	}
	app.Action = func(c *cli.Context) error {
		enabledMetrics := make([]string, 0)
		metrics := make([]Metric, 0)
		if c.Bool("memory") {
			metrics = append(metrics, Memory{})
			enabledMetrics = append(enabledMetrics, "memory")
		}
		if c.Bool("swap") {
			metrics = append(metrics, Swap{})
			enabledMetrics = append(enabledMetrics, "swap")
		}
		if c.Bool("disk") {
			metrics = append(metrics, Disk{})
			enabledMetrics = append(enabledMetrics, "disk")
		}
		if c.Bool("network") {
			metrics = append(metrics, Network{})
			enabledMetrics = append(enabledMetrics, "network")
		}
		if c.Bool("docker") {
			metrics = append(metrics, Docker{})
			enabledMetrics = append(enabledMetrics, "docker")
		}

		cfg, err := external.LoadDefaultAWSConfig()
		if err != nil {
			panic("Unable to load SDK config")
		}

		cfg.Region = c.String("region")
		cloudWatch := CloudWatchService{
			Config: cfg,
		}

		interval := c.Int("interval")

		fmt.Printf("Interval: %d minutes - Features enabled: %s\n", interval, strings.Join(enabledMetrics, ", "))

		duration := time.Duration(interval) * time.Minute
		for _ = range time.Tick(duration) {
			Collect(metrics, cloudWatch)
		}

		return nil
	}
	app.Run(os.Args)
}
