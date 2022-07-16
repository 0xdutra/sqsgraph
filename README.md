## SQS Graph

View metrics for AWS SQS queues in your terminal.

![](./img/graph.png)

<hr>

## Flags

```
  -interval int
    	Internal to updating graph (default 10)
  -list-queues
    	List all SQS queues
  -metric string
    	The metric name to visualize in the graph, default is ApproximateNumberOfMessages (default "ApproximateNumberOfMessages")
  -queue-url string
    	The url of the SQS queue
```

## Build

```sh
go build
```

## Example

```sh
sqsgraph -queue <queue url>
```
