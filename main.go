package main

import "fmt"

const (
	sqscToken           = "SQSC_TOKEN"
	dockerRepository    = "DOCKER_REPOSITORY"
	dockerRepositoryTag = "DOCKER_REPOSITORY_TAG"
	organizationName    = "ORGANIZATION_NAME"
	projectName         = "PROJECT_NAME"
	iaasProvider        = "IAAS_PROVIDER"
	iaasRegion          = "IAAS_REGION"
	iaasCred            = "IAAS_CRED"
	monitoring          = "MONITORING"
	infraType           = "INFRA_TYPE"
	nodeType            = "NODE_TYPE"
	dbEngine            = "DB_ENGINE"
	dbEngineVersion     = "DB_ENGINE_VERSION"
	dbSize              = "DB_SIZE"
	servicesEnv         = "SERVICES"
	batchesEnv          = "BATCHES"
)

func main() {
	checkEnvVarError := checkEnvironmentVariablesExists()
	if checkEnvVarError != nil {
		fmt.Println(checkEnvVarError)
	}

	project := Project{}
	project.create()

	database := Database{}
	err := database.create()
	if err != nil {
		fmt.Println(err)
	}

	services := Services{}
	services.create()

	batches := Batches{}
	batches.create()
}
