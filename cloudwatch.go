package main

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type CloudWatch struct {
	cfg aws.Config
}

func (c CloudWatch) Publish(metricData []cloudwatch.MetricDatum, namespace string) {
	svc := cloudwatch.New(c.cfg)
	req := svc.PutMetricDataRequest(&cloudwatch.PutMetricDataInput{
		MetricData: metricData,
		Namespace:  &namespace,
	})
	_, err := req.Send()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Metric %s has been save\n", namespace)
	}
}
