package test

import (
	"os"
	"testing"

	"eventlogger/config"
)

func TestGetEnvWithFallback(t *testing.T) {
	os.Unsetenv("NOT_SET_VAR")
	result := config.GetEnv("NOT_SET_VAR", "fallback")
	if result != "fallback" {
		t.Errorf("esperaba 'fallback', obtuvo '%s'", result)
	}
}
