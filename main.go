package main

import (
	"flag"
	"os"
	"sqsgraph/src/graph"
	"sqsgraph/src/sqs"
)

var (
	queueUrl        *string
	refreshInterval *int64
	metricName      *string
	listQueues      *bool
)

func init() {
	queueUrl = flag.String("queue-url", "", "The url of the SQS queue")
	refreshInterval = flag.Int64("interval", 10, "Internal to updating graph")
	metricName = flag.String("metric", "ApproximateNumberOfMessages", "The metric name to visualize in the graph, default is ApproximateNumberOfMessages")
	listQueues = flag.Bool("list-queues", false, "List all SQS queues")
}

func main() {
	flag.Parse()

	if *listQueues {
		sqs.ListAllQueues()
		os.Exit(0)
	}

	graph.Graph(*queueUrl, *metricName, *refreshInterval)
}
