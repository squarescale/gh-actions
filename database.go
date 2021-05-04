package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/hoisie/mustache"
)

type Database struct{}

func (d *Database) create() error {
	_, dbEngineExists := os.LookupEnv(dbEngine)
	_, dbEngineVersionExists := os.LookupEnv(dbEngineVersion)
	_, dbEngineSizeExists := os.LookupEnv(dbSize)

	if !dbEngineExists && !dbEngineVersionExists && !dbEngineSizeExists {
		return errors.New(fmt.Sprintf("%s, %s, %s are not set correctly. No database will be created.", dbEngine, dbEngineVersion, dbSize))
	}

	if !isDatabaseExists() {
		d.createDatabase()
	} else {
		fmt.Println("Database already exists.")
	}

	return nil
}

func (d *Database) createDatabase() string {
	fmt.Println("Creating database...")

	cmd := fmt.Sprintf(
		"/sqsc db set -project-name %s -engine \"%s\" -engine-version \"%s\" -size \"%s\" -yes",
		getProjectName(),
		os.Getenv(dbEngine),
		os.Getenv(dbEngineVersion),
		os.Getenv(dbSize),
	)
	executeCommand(cmd, "Fail to create database.")
	return cmd
}

func isDatabaseExists() bool {
	databaseNotExists := executeCommand(fmt.Sprintf(
		"/sqsc db show -project-name %s | grep \"DB enabled\" | grep true",
		getProjectName(),
	), "Fail to check if database exists.")

	return databaseNotExists == nil
}

func mapDatabaseEnv(env string) string {
	return mustache.Render(env, map[string]string{
		"DB_HOST":     getSQSCEnvValue("DB_HOST"),
		"DB_USERNAME": getSQSCEnvValue("DB_USERNAME"),
		"DB_PASSWORD": getSQSCEnvValue("DB_PASSWORD"),
		"DB_NAME":     getSQSCEnvValue("DB_NAME"),
	})
}
