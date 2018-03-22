package services

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

type CloudWatchService struct {
	Config aws.Config
}

func (c CloudWatchService) Publish(metricData []cloudwatch.MetricDatum, namespace string) {
	svc := cloudwatch.New(c.Config)
	req := svc.PutMetricDataRequest(&cloudwatch.PutMetricDataInput{
		MetricData: metricData,
		Namespace:  &namespace,
	})
	_, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
}
