package main

import (
	"fmt"
	"os"
)

type Project struct{}

func (p *Project) create() {
	if !isProjectExists() {
		p.createProject()
	} else {
		fmt.Println("Project already exists.")
	}
}

func (p *Project) createProject() string {
	fmt.Println("Creating project...")

	cmd := "/sqsc project create"
	cmd += " -credential " + os.Getenv(iaasCred)
	cmd += " -name " + os.Getenv(projectName)
	cmd += " -node-size " + os.Getenv(nodeType)
	cmd += " -infra-type " + os.Getenv(infraType)
	cmd += " -organization " + os.Getenv(organizationName)
	cmd += " -provider " + os.Getenv(iaasProvider)
	cmd += " -region " + os.Getenv(iaasRegion)

	if os.Getenv(monitoring) != "" {
		cmd += " -monitoring " + os.Getenv(monitoring)
	}

	cmd += " -yes"

	executeCommand(cmd, fmt.Sprintf("Fail to create project %q.", os.Getenv(projectName)))
	return cmd
}

func isProjectExists() bool {
	projectNotExists := executeCommand(fmt.Sprintf(
		"/sqsc project get -project-name %s",
		getProjectName(),
	), "Fail to check if project exists.")

	return projectNotExists == nil
}
