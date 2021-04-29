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
}

func TestCreateBatchWithEmptyBatchContentAndBatchName(t *testing.T) {
	// given
	batches := Batches{}
	batchContent := BatchContent{}
	expectedCMD := "/sqsc batch add -name  -project-name  -run-command '' -imageName "
	executeCommand = func(cmd string, errorMsg string) {}

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
	executeCommand = func(cmd string, errorMsg string) {}

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
	executeCommand = func(cmd string, errorMsg string) {}

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
	executeCommand = func(cmd string, errorMsg string) {}

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
	executeCommand = func(cmd string, errorMsg string) {}

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
	executeCommand = func(cmd string, errorMsg string) {}

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
	executeCommand = func(cmd string, errorMsg string) {}

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
	executeCommand = func(cmd string, errorMsg string) {}

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
	executeCommand = func(cmd string, errorMsg string) {}

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
	executeCommand = func(cmd string, errorMsg string) {}

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
	executeCommand = func(cmd string, errorMsg string) {}
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
	executeCommand = func(cmd string, errorMsg string) {}
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
	executeCommand = func(cmd string, errorMsg string) {}
	expectedCMD := "/sqsc batch add -name  -project-name  -run-command '' -imageName  -periodic -cron '1 * * * *' -time \"Europe/Test\""
	// when
	cmd, _ := batches.createBatch("", batchContent)

	// then
	if expectedCMD != cmd {
		t.Fatalf("Command not correct, expected to be '%s', but got '%s'", expectedCMD, cmd)
	}
}
