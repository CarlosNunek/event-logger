package test

import (
	"testing"
	"time"

	"eventlogger/influx"
)

func TestWriteEvent(t *testing.T) {
	err := influx.WriteEvent("test-canal", "mensaje de prueba", time.Now())
	if err != nil {
		t.Errorf("falló al escribir en Influx: %v", err)
	}
}
