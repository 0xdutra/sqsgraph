package graph

import (
	"fmt"
	"sqsmonit/src/sqs"
	"time"

	"github.com/guptarohit/asciigraph"
)

var (
	height       uint = 20
	width        uint = 40
	offset       uint = 3
	precision    uint = 2
	caption      string
	fps          float64 = 24
	seriesColor  asciigraph.AnsiColor
	captionColor asciigraph.AnsiColor
	axisColor    asciigraph.AnsiColor
	labelColor   asciigraph.AnsiColor
)

func Graph(queueUrl string, metricName string, refreshInterval int64) {
	var data []float64

	nextFlushTime := time.Now()
	flushInterval := time.Duration(float64(time.Second) / fps)

	for {
		metricValue := sqs.QueueAttributes(queueUrl, metricName)
		data = append(data, float64(metricValue))

		if currentTime := time.Now(); currentTime.After(nextFlushTime) || currentTime.Equal(nextFlushTime) {
			plot := asciigraph.Plot(data,
				asciigraph.Height(int(height)),
				asciigraph.Width(int(width)),
				asciigraph.Offset(int(offset)),
				asciigraph.Precision(precision),
				asciigraph.Caption(caption),
				asciigraph.SeriesColors(seriesColor),
				asciigraph.CaptionColor(captionColor),
				asciigraph.AxisColor(axisColor),
				asciigraph.LabelColor(labelColor),
			)
			asciigraph.Clear()
			fmt.Println(plot)
			fmt.Printf("\nQueue URL: %s\n", queueUrl)
			fmt.Printf("Current value: %d\n", metricValue)

			nextFlushTime = time.Now().Add(flushInterval)
		}

		time.Sleep(time.Duration(refreshInterval) * time.Second)
	}
}
