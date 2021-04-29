package main

import (
	"os"
	"testing"
)

func setTestEnv(t *testing.T, env map[string]string) {
	os.Clearenv()
	for key, value := range env {
		err := os.Setenv(key, value)
		if err != nil {
			t.Fatalf("Error when setting env `%s` whith value `%s`", key, value)
		}
	}
}
