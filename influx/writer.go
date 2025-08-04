package influx

import (
	"context"
	"time"

	"eventlogger/config"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var writeAPI api.WriteAPIBlocking

func init() {
	url := config.GetEnv("INFLUX_URL", "http://localhost:8086")
	token := config.GetEnv("INFLUX_TOKEN", "ICusY4BPCbyxTKIxAtp7ciEalt96codUyofHs6wzANAmAH0TX9vTOGzMXor3Ryzgys-gfAHCUszY1WOcNz4v5A==")
	org := config.GetEnv("INFLUX_ORG", "my-org")
	bucket := config.GetEnv("INFLUX_BUCKET", "eventos")

	client := influxdb2.NewClient(url, token)
	writeAPI = client.WriteAPIBlocking(org, bucket)
}

func WriteEvent(canal, payload string, t time.Time) error {
	p := influxdb2.NewPointWithMeasurement("eventos").
		AddTag("canal", canal).
		AddField("mensaje", payload).
		SetTime(t)
	return writeAPI.WritePoint(context.Background(), p)
}
