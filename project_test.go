package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProject(t *testing.T) {
	t.Run("Create project if not exists", CreateProjectIfNotExists)
	t.Run("Not create project if exists", NotCreateProjectIfExists)
	t.Run("Create project without monitoring", CreateProjectWithoutMonitoring)
	t.Run("Create project with monitoring", CreateProjectWithMonitoring)
}

func CreateProjectIfNotExists(t *testing.T) {
	// given
	project := Project{}
	executeCommand = func(cmd string, errorMsg string) interface{} { return true }

	// when
	project.create()
}

func NotCreateProjectIfExists(t *testing.T) {
	// given
	project := Project{}
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	project.create()
}

func CreateProjectWithoutMonitoring(t *testing.T) {
	// given
	project := Project{}
	env := map[string]string{
		iaasCred:         "iaasCred",
		projectName:      "projectName",
		nodeType:         "nodeType",
		infraType:        "infraType",
		organizationName: "organizationName",
		iaasProvider:     "iaasProvider",
		iaasRegion:       "iaasRegion",
	}
	setTestEnv(t, env)
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	cmd := project.createProject()

	// then
	expectedCmd := "/sqsc project create -credential iaasCred -name projectName -node-size nodeType -infra-type infraType -organization organizationName -provider iaasProvider -region iaasRegion -yes"
	assert.Equal(t, expectedCmd, cmd)
}

func CreateProjectWithMonitoring(t *testing.T) {
	// given
	project := Project{}
	env := map[string]string{
		iaasCred:         "iaasCred",
		projectName:      "projectName",
		nodeType:         "nodeType",
		infraType:        "infraType",
		organizationName: "organizationName",
		iaasProvider:     "iaasProvider",
		iaasRegion:       "iaasRegion",
		monitoring:       "monitoring",
	}
	setTestEnv(t, env)
	executeCommand = func(cmd string, errorMsg string) interface{} { return nil }

	// when
	cmd := project.createProject()

	// then
	expectedCmd := "/sqsc project create -credential iaasCred -name projectName -node-size nodeType -infra-type infraType -organization organizationName -provider iaasProvider -region iaasRegion -monitoring monitoring -yes"
	assert.Equal(t, expectedCmd, cmd)
}
