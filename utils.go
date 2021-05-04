package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var executeCommand = executeSQSCCommand

func checkEnvironmentVariablesExists() error {
	fmt.Println("Checking environment variables...")

	envVars := []string{
		sqscToken,
		dockerRepository,
		projectName,
		iaasProvider,
		iaasRegion,
		iaasCred,
		nodeType,
		infraType,
	}

	for _, envVar := range envVars {
		if _, exists := os.LookupEnv(envVar); !exists {
			os.Exit(1)
			return errors.New(fmt.Sprintf("%s is not set. Quitting.", envVar))
		}
	}

	return nil
}

func executeSQSCCommand(cmd string, errorMsg string) interface{} {
	fmt.Println(cmd)
	output, err := exec.Command("/bin/sh", "-c", cmd).Output()

	if err != nil {
		fmt.Println(string(output))
		log.Fatal(errorMsg)
	}

	fmt.Println(string(output))
	return output
}

func getProjectName() string {
	if _, exists := os.LookupEnv(organizationName); exists {
		return fmt.Sprintf("%s/%s", os.Getenv(organizationName), os.Getenv(projectName))
	}

	return fmt.Sprintf("%s", os.Getenv(projectName))
}

func getDockerImage() string {
	if _, exists := os.LookupEnv(dockerRepositoryTag); exists {
		return fmt.Sprintf("%s:%s", os.Getenv(dockerRepository), os.Getenv(dockerRepositoryTag))
	}

	return fmt.Sprintf("%s", os.Getenv(dockerRepository))
}

func getSQSCEnvValue(key string) string {
	value, err := exec.Command("/bin/sh", "-c", fmt.Sprintf(
		"/sqsc env get -project-name %s \"%s\" | grep -v %s | tr -d '\n'",
		getProjectName(),
		key,
		"...done",
	)).Output()

	if err != nil {
		fmt.Println(fmt.Sprintf("Environment variable %q does not exists in this project.", key))
		return ""
	}

	return string(value)
}

func isNetworkRuleExists(networkRuleName string, serviceName string) bool {
	_, networkRuleNotExists := exec.Command("/bin/sh", "-c", fmt.Sprintf(
		"/sqsc network-rule list -project-name %s -service-name %s | grep %s",
		getProjectName(),
		serviceName,
		networkRuleName,
	)).Output()

	return networkRuleNotExists == nil
}
