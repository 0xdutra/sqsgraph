package sqs

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ListAllQueues() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	result, err := svc.ListQueues(&sqs.ListQueuesInput{})
	if err != nil {
		log.Panic(err)
	}

	for _, queue := range result.QueueUrls {
		fmt.Println(*queue)
	}
}

func QueueAttributes(queueUrl string, metricName string) uint64 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	result, err := svc.GetQueueAttributes(&sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(queueUrl),
		AttributeNames: []*string{
			aws.String(metricName),
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	var metricValue string

	for _, attr := range result.Attributes {
		metricValue = *attr
	}

	metricValueInt, err := strconv.ParseUint(metricValue, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	return metricValueInt
}
