package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {
	t.Run("Check environment variables exists", CheckEnvironementVariablesExists)
}

func CheckEnvironementVariablesExists(t *testing.T) {
	// given
	env := map[string]string{
		sqscToken:        "sqscToken",
		dockerRepository: "dockerRepository",
		projectName:      "projectName",
		iaasProvider:     "iaasProvider",
		iaasRegion:       "iaasRegion",
		iaasCred:         "iaasCred",
		nodeType:         "nodeType",
		infraType:        "infraType",
	}
	setTestEnv(t, env)

	// when
	err := checkEnvironmentVariablesExists()

	// then
	assert.NoError(t, err)
}

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
