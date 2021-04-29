package main

import (
	"encoding/json"
	"io/ioutil"
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

func readEnvFile() (map[string]interface{}, error) {
	jsonFile, err := os.Open("batchEnvVar.json")
	byteValue, err := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	return result, err
}
