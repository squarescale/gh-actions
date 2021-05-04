package main

import (
	"os"
	"testing"
)

func TestBatches(t *testing.T) {
	t.Run("Create batch with empty batch content and empty batch name", TestCreateBatchWithEmptyBatchContentAndBatchName)
	t.Run("Create batch with batch name", TestCreateBatchWithBatchName)
	t.Run("Create batch with project name", TestCreateBatchWithProjectName)
	t.Run("Create batch with run command", TestCreateBatchWithRunCMD)
	t.Run("Create batch with image name", TestCreateBatchWithImageName)
	t.Run("Create batch with env docker repository tag", TestCreateBatchWithEnvDockerRepositoryTag)
	t.Run("Create batch with env docker repository tag and image name", TestCreateBatchWithEnvDockerRepositoryTagAndImageName)
	t.Run("Create batch with private registry", TestCreateBatchWithIsPrivate)
	t.Run("Create batch with private registry and empty user", TestCreateBatchWithIsPrivateAndEmptyUser)
	t.Run("Create batch with private registry and empty password", TestCreateBatchWithIsPrivateAndEmptyPassword)
	t.Run("Create periodic batch with periodicity", TestCreatePeriodicBatchWithPeriodicity)
	t.Run("Create periodic batch with time zone", TestCreatePeriodicBatchWithTimeZone)
	t.Run("Create periodic batch", TestCreatePeriodicBatch)

	t.Run("Insert batch env and limits with empty values", TestInsertBatchEnvAndLimitsWithEmptyValues)
	t.Run("Insert batch env and limits with memory limit", TestInsertBatchEnvAndLimitsWithMemoryLimit)
	t.Run("Insert batch env and limits with net limit", TestInsertBatchEnvAndLimitsWithNetLimit)
	t.Run("Insert batch env and limits with CPU limit", TestInsertBatchEnvAndLimitsWithCPULimit)
	t.Run("Insert batch env and limits with env", TestInsertBatchEnvAndLimitsWithEnv)
}

func TestCreateBatchWithEmptyBatchContentAndBatchName(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{}
	expectedCMD := "/sqsc batch add -name  -project-name  -run-command '' -imageName "
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	cmd, _ := batches.createBatch("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestCreateBatchWithBatchName(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{}
	expectedCMD := "/sqsc batch add -name Test-batch -project-name  -run-command '' -imageName "
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	cmd, _ := batches.createBatch("Test-batch", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestCreateBatchWithProjectName(t *testing.T) {
	// given
	env := map[string]string{
		organizationName: "orga-test",
		projectName:      "project-test",
	}
	setTestEnv(t, env)
	defer os.Clearenv()
	batches := Batches{}
	batchContent := BatchContent{}
	expectedCMD := "/sqsc batch add -name  -project-name orga-test/project-test -run-command '' -imageName "
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	cmd, _ := batches.createBatch("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestCreateBatchWithRunCMD(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{
		RUN_CMD: "Test command...",
	}
	expectedCMD := "/sqsc batch add -name  -project-name  -run-command 'Test command...' -imageName "
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	cmd, _ := batches.createBatch("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestCreateBatchWithImageName(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{
		IMAGE_NAME: "TestImage",
	}
	expectedCMD := "/sqsc batch add -name  -project-name  -run-command '' -imageName TestImage"
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	cmd, _ := batches.createBatch("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestCreateBatchWithEnvDockerRepositoryTag(t *testing.T) {
	// given
	env := map[string]string{
		dockerRepository:    "dockerTest",
		dockerRepositoryTag: "tagTest",
	}
	setTestEnv(t, env)
	defer os.Clearenv()
	batches := Batches{}
	batchContent := BatchContent{}
	expectedCMD := "/sqsc batch add -name  -project-name  -run-command '' -imageName dockerTest:tagTest"
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	cmd, _ := batches.createBatch("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestCreateBatchWithEnvDockerRepositoryTagAndImageName(t *testing.T) {
	// given
	env := map[string]string{
		dockerRepository:    "dockerTest",
		dockerRepositoryTag: "tagTest",
	}
	setTestEnv(t, env)
	defer os.Clearenv()
	batches := Batches{}
	batchContent := BatchContent{
		IMAGE_NAME: "TestImage",
	}
	expectedCMD := "/sqsc batch add -name  -project-name  -run-command '' -imageName TestImage"
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	cmd, _ := batches.createBatch("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestCreateBatchWithIsPrivate(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{
		IS_PRIVATE:     true,
		IMAGE_USER:     "userTest",
		IMAGE_PASSWORD: "12345",
	}
	expectedCMD := "/sqsc batch add -name  -project-name  -run-command '' -imageName  -imagePrivate -imageUser userTest -imagePwd 12345"
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	cmd, _ := batches.createBatch("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestCreateBatchWithIsPrivateAndEmptyUser(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{
		IS_PRIVATE:     true,
		IMAGE_PASSWORD: "12345",
	}
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	_, err := batches.createBatch("", batchContent)

	// then
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestCreateBatchWithIsPrivateAndEmptyPassword(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{
		IS_PRIVATE: true,
		IMAGE_USER: "userTest",
	}
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	_, err := batches.createBatch("", batchContent)

	// then
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestCreatePeriodicBatchWithPeriodicity(t *testing.T) {
	// given
	batches := Batches{}
	periodicBatch := BatchPeriodic{
		PERIODICITY: "1 * * * *",
	}
	batchContent := BatchContent{
		PERIODIC: periodicBatch,
	}
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }
	expectedCMD := "/sqsc batch add -name  -project-name  -run-command '' -imageName  -periodic -cron '1 * * * *' -time \"Europe/Paris\""
	// when
	cmd, _ := batches.createBatch("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestCreatePeriodicBatchWithTimeZone(t *testing.T) {
	// given
	batches := Batches{}
	periodicBatch := BatchPeriodic{
		TIMEZONE: "Europe/Test",
	}
	batchContent := BatchContent{
		PERIODIC: periodicBatch,
	}
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }
	expectedCMD := "/sqsc batch add -name  -project-name  -run-command '' -imageName  -periodic -cron '* * * * *' -time \"Europe/Test\""
	// when
	cmd, _ := batches.createBatch("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestCreatePeriodicBatch(t *testing.T) {
	// given
	batches := Batches{}
	periodicBatch := BatchPeriodic{
		TIMEZONE:    "Europe/Test",
		PERIODICITY: "1 * * * *",
	}
	batchContent := BatchContent{
		PERIODIC: periodicBatch,
	}
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }
	expectedCMD := "/sqsc batch add -name  -project-name  -run-command '' -imageName  -periodic -cron '1 * * * *' -time \"Europe/Test\""
	// when
	cmd, _ := batches.createBatch("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestInsertBatchEnvAndLimitsWithEmptyValues(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{}
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }
	expectedCMD := ""
	// when
	cmd := batches.insertBatchEnvAndLimits("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestInsertBatchEnvAndLimitsWithCPULimit(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{
		LIMIT_CPU: "12",
	}
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }
	expectedCMD := "/sqsc batch set -project-name  -batch-name testBatch -cpu 12"
	// when
	cmd := batches.insertBatchEnvAndLimits("testBatch", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestInsertBatchEnvAndLimitsWithMemoryLimit(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{
		LIMIT_MEMORY: "1200",
	}
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }
	expectedCMD := "/sqsc batch set -project-name  -batch-name testBatch -memory 1200"
	// when
	cmd := batches.insertBatchEnvAndLimits("testBatch", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestInsertBatchEnvAndLimitsWithNetLimit(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{
		LIMIT_NET: "12",
	}
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }
	expectedCMD := "/sqsc batch set -project-name  -batch-name testBatch -net 12"
	// when
	cmd := batches.insertBatchEnvAndLimits("testBatch", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}

func TestInsertBatchEnvAndLimitsWithEnv(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{
		ENV: map[string]string{
			"DB_HOST": "test.host",
			"DB_NAME": "dbTest",
		},
	}
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }
	expectedCMD := "/sqsc batch set -project-name  -batch-name testBatch -env batchEnvVar.json"
	// when
	cmd := batches.insertBatchEnvAndLimits("testBatch", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}

	mapEnv, err := readEnvFile()

	if err != nil {
		t.Fatalf("Expected no error, got '%s'", err)
	}

	for key, value := range mapEnv {
		switch key {
		case "DB_HOST":
			expectedEnv := "test.host"
			if value != expectedEnv {
				t.Fatalf("Env not correct, expected to be '%s', but got '%s'", expectedEnv, value)
			}
		case "DB_NAME":
			expectedEnv := "dbTest"
			if value != expectedEnv {
				t.Fatalf("Env not correct, expected to be '%s', but got '%s'", expectedEnv, value)
			}
		default:
			t.Fatalf("Env key '%s' is not correct", key)
		}
	}
}
