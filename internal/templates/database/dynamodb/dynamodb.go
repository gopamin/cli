package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var dynamodbEnv []byte

func DynamodbEnvTemplate() ([]byte, string) {
	return dynamodbEnv, ".env"
}

//go:embed files/makefile.tmpl
var dynamodbMakefile []byte

func DynamodbMakefileTemplate() ([]byte, string) {
	return dynamodbMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var dynamodbReadme []byte

func DynamodbReadmeTemplate() ([]byte, string) {
	return dynamodbReadme, "README.md"
}

//go:embed files/docker-compose.tmpl
var dynamodbDockerCompose []byte

func DynamodbDockerComposeTemplate() ([]byte, string) {
	return dynamodbDockerCompose, "docker-compose.yml"
}

//go:embed files/internal/adapters/repositories/dynamodb/repository.tmpl
var dynamodbRepository []byte

func DynamodbRepositoryTemplate() ([]byte, string) {
	return dynamodbRepository, "internal/adapters/repositories/dynamodb/repository.go"
}
