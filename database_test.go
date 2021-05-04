package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	t.Run("Create database if not exists", CreateDatabaseIfNotExists)
	t.Run("Do not create database if exists", DoNotCreateDatabaseIfExists)
	t.Run("Create database with no environement variables", CreateDatabaseWithNoEnvVar)
	t.Run("Create database", CreateDatabase)
	t.Run("Check if database exists", CheckIfDatabaseExists)
}

func CreateDatabaseIfNotExists(t *testing.T) {
	// given
	database := Database{}
	env := map[string]string{
		projectName:     "test",
		dbEngine:        "postgres",
		dbEngineVersion: "12",
		dbSize:          "small",
	}
	setTestEnv(t, env)
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	err := database.create()

	// then
	assert.NoError(t, err)
}

func DoNotCreateDatabaseIfExists(t *testing.T) {
	// given
	database := Database{}
	env := map[string]string{
		projectName:     "test",
		dbEngine:        "postgres",
		dbEngineVersion: "12",
		dbSize:          "small",
	}
	setTestEnv(t, env)
	executeCommand = func(cmd string, errorMsg string) interface{} { return true }

	// when
	err := database.create()

	// then
	assert.NoError(t, err)
}

func CreateDatabaseWithNoEnvVar(t *testing.T) {
	// given
	database := Database{}
	os.Clearenv()
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	err := database.create()

	// then
	assert.Error(t, err)
}

func CreateDatabase(t *testing.T) {
	// given
	database := Database{}
	env := map[string]string{
		projectName:     "test",
		dbEngine:        "postgres",
		dbEngineVersion: "12",
		dbSize:          "small",
	}
	setTestEnv(t, env)
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	cmd := database.createDatabase()

	// then
	expectedCMD := "/sqsc db set -project-name test -engine \"postgres\" -engine-version \"12\" -size \"small\" -yes"
	assert.Equal(t, expectedCMD, cmd)
}

func CheckIfDatabaseExists(t *testing.T) {
	// given
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	exists := isDatabaseExists()

	// then
	assert.True(t, exists)
}
