package main

import (
	"flag"
	"sqsmonit/src/graph"
)

var (
	queueUrl        *string
	refreshInterval *int64
	metricName      *string
)

func init() {
	queueUrl = flag.String("queue", "", "The url of the SQS queue")
	refreshInterval = flag.Int64("interval", 10, "Internal to updating graph")
	metricName = flag.String("metric", "ApproximateNumberOfMessages", "The metric name to visualize in the graph, default is ApproximateNumberOfMessages")
}

func main() {
	flag.Parse()
	graph.Graph(*queueUrl, *metricName, *refreshInterval)
}
