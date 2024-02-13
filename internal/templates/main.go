package templates

import (
	apiChi "github.com/gopamin/cli/internal/templates/api/chi"
	apiEcho "github.com/gopamin/cli/internal/templates/api/echo"
	apiGin "github.com/gopamin/cli/internal/templates/api/gin"
	apiGorilla "github.com/gopamin/cli/internal/templates/api/gorilla"
	apiHttp "github.com/gopamin/cli/internal/templates/api/http"
	apiHttprouter "github.com/gopamin/cli/internal/templates/api/httprouter"
	common "github.com/gopamin/cli/internal/templates/common"
	dynamodb "github.com/gopamin/cli/internal/templates/database/dynamodb"
	mongodb "github.com/gopamin/cli/internal/templates/database/mongodb"
	mysql "github.com/gopamin/cli/internal/templates/database/mysql"
	postgres "github.com/gopamin/cli/internal/templates/database/postgres"
	redis "github.com/gopamin/cli/internal/templates/database/redis"
	sqlite "github.com/gopamin/cli/internal/templates/database/sqlite"
	helloWorld "github.com/gopamin/cli/internal/templates/hello-world"
	kafkaMicroservice "github.com/gopamin/cli/internal/templates/microservice/kafka"
	rabbitmqMicroservice "github.com/gopamin/cli/internal/templates/microservice/rabbitmq"
	redisMicroservice "github.com/gopamin/cli/internal/templates/microservice/redis"
	webApp "github.com/gopamin/cli/internal/templates/web-app"
)

func Mapper() map[string]func() ([]byte, string) {
	return map[string]func() ([]byte, string){
		"dockerfile":                common.DockerFileTemplate,
		"gitignore":                 common.GitIgnoreTemplate,
		"license":                   common.LicenseTemplate,
		"dockerignore":              common.DockerIgnoreTemplate,
		"mock-repository":           common.MockRepositoryTemplate,
		"user-test":                 common.UserTestTemplate,
		"user":                      common.UserTemplate,
		"user-repository-interface": common.UserRepositoryInterfaceTemplate,
		"user-service-interface":    common.UserServiceInterfaceTemplate,
		"user-service":              common.UserServiceTemplate,
		"user-service-test":         common.UserServiceTestTemplate,
		"router-inferface":          common.RouterInterfaceTemplate,
		"api-errors":                common.ApiErrorsTemplate,
		"api-response":              common.ApiResponseTemplate,
		"load-env":                  common.LoadEnvTemplate,
		"api-server":                common.ApiServerTemplate,
		"api-main":                  common.ApiMainTemplate,
		"api-main-with-db":          common.ApiMainWithDbTemplate,
		"hello-world-main":          common.HelloWorldMainTemplate,
		"hello-world-main-with-db":  common.HelloWorldMainWithDbTemplate,
		"web-app-index-page":        common.WebAppIndexPageTemplate,
		"web-app-styles":            common.WebAppIndexPageTemplate,
		"web-app-main":              common.WebAppMainTemplate,
		"message":                   common.MessageTemplate,
		"broker-service-interface":  common.BrokerServiceInterfaceTemplate,
		"broker-service":            common.BrokerServiceTemplate,
		"message-broker-interface":  common.MessageBrokerInterfaceTemplate,
		"microservice-main":         common.MicroserviceMainTemplate,

		"hello-world-readme":   helloWorld.HelloWorldReadmeTemplate,
		"hello-world-makefile": helloWorld.HelloWorldMakefileTemplate,
		"hello-world-env":      helloWorld.HelloWorldEnvTemplate,

		"api-http-readme":   apiHttp.ApiHttpReadmeTemplate,
		"api-http-makefile": apiHttp.ApiHttpMakefileTemplate,
		"api-http-env":      apiHttp.ApiHttpEnvTemplate,
		"api-http-routes":   apiHttp.ApiHttpRoutesTemplate,
		"api-http-users":    apiHttp.ApiHttpUsersTemplate,

		"api-chi-readme":   apiChi.ApiChiReadmeTemplate,
		"api-chi-makefile": apiChi.ApiChiMakefileTemplate,
		"api-chi-env":      apiChi.ApiChiEnvTemplate,
		"api-chi-routes":   apiChi.ApiChiRoutesTemplate,
		"api-chi-users":    apiChi.ApiChiUsersTemplate,

		"api-echo-readme":   apiEcho.ApiEchoReadmeTemplate,
		"api-echo-makefile": apiEcho.ApiEchoMakefileTemplate,
		"api-echo-env":      apiEcho.ApiEchoEnvTemplate,
		"api-echo-routes":   apiEcho.ApiEchoRoutesTemplate,
		"api-echo-users":    apiEcho.ApiEchoUsersTemplate,

		"api-gin-readme":   apiGin.ApiGinReadmeTemplate,
		"api-gin-makefile": apiGin.ApiGinMakefileTemplate,
		"api-gin-env":      apiGin.ApiGinEnvTemplate,
		"api-gin-routes":   apiGin.ApiGinRoutesTemplate,
		"api-gin-users":    apiGin.ApiGinUsersTemplate,

		"api-gorilla-readme":   apiGorilla.ApiGorillaReadmeTemplate,
		"api-gorilla-makefile": apiGorilla.ApiGorillaMakefileTemplate,
		"api-gorilla-env":      apiGorilla.ApiGorillaEnvTemplate,
		"api-gorilla-routes":   apiGorilla.ApiGorillaRoutesTemplate,
		"api-gorilla-users":    apiGorilla.ApiGorillaUsersTemplate,

		"api-httprouter-readme":   apiHttprouter.ApiHttprouterReadmeTemplate,
		"api-httprouter-makefile": apiHttprouter.ApiHttprouterMakefileTemplate,
		"api-httprouter-env":      apiHttprouter.ApiHttprouterEnvTemplate,
		"api-httprouter-routes":   apiHttprouter.ApiHttprouterRoutesTemplate,
		"api-httprouter-users":    apiHttprouter.ApiHttprouterUsersTemplate,

		"mysql-env":            mysql.MysqlEnvTemplate,
		"mysql-repository":     mysql.MysqlRepositoryTemplate,
		"mysql-readme":         mysql.MysqlReadmeTemplate,
		"mysql-makefile":       mysql.MysqlMakefileTemplate,
		"mysql-docker-compose": mysql.MysqlDockerComposeTemplate,

		"postgres-env":            postgres.PostgresEnvTemplate,
		"postgres-repository":     postgres.PostgresRepositoryTemplate,
		"postgres-readme":         postgres.PostgresReadmeTemplate,
		"postgres-makefile":       postgres.PostgresMakefileTemplate,
		"postgres-docker-compose": postgres.PostgresDockerComposeTemplate,

		"mongodb-env":            mongodb.MongodbEnvTemplate,
		"mongodb-repository":     mongodb.MongodbRepositoryTemplate,
		"mongodb-readme":         mongodb.MongodbReadmeTemplate,
		"mongodb-makefile":       mongodb.MongodbMakefileTemplate,
		"mongodb-docker-compose": mongodb.MongodbDockerComposeTemplate,

		"redis-env":            redis.RedisEnvTemplate,
		"redis-repository":     redis.RedisRepositoryTemplate,
		"redis-readme":         redis.RedisReadmeTemplate,
		"redis-makefile":       redis.RedisMakefileTemplate,
		"redis-docker-compose": redis.RedisDockerComposeTemplate,

		"sqlite-env":        sqlite.SqliteEnvTemplate,
		"sqlite-repository": sqlite.SqliteRepositoryTemplate,
		"sqlite-makefile":   sqlite.SqliteMakefileTemplate,
		"sqlite-readme":     sqlite.SqliteReadmeTemplate,

		"dynamodb-env":            dynamodb.DynamodbEnvTemplate,
		"dynamodb-repository":     dynamodb.DynamodbRepositoryTemplate,
		"dynamodb-readme":         dynamodb.DynamodbReadmeTemplate,
		"dynamodb-makefile":       dynamodb.DynamodbMakefileTemplate,
		"dynamodb-docker-compose": dynamodb.DynamodbDockerComposeTemplate,

		"web-app-readme":   webApp.WebAppReadmeTemplate,
		"web-app-env":      webApp.WebAppEnvTemplate,
		"web-app-makefile": webApp.WebAppMakefileTemplate,

		"redis-microservice-readme":         redisMicroservice.RedisMicroserviceReadmeTemplate,
		"redis-microservice-env":            redisMicroservice.RedisMicroserviceEnvTemplate,
		"redis-microservice-makefile":       redisMicroservice.RedisMicroserviceMakefileTemplate,
		"redis-microservice-docker-compose": redisMicroservice.RedisMicroserviceDockerComposeTemplate,
		"redis-microservice-broker":         redisMicroservice.RedisMicroserviceBrokerTemplate,

		"kafka-microservice-readme":         kafkaMicroservice.KafkaMicroserviceReadmeTemplate,
		"kafka-microservice-env":            kafkaMicroservice.KafkaMicroserviceEnvTemplate,
		"kafka-microservice-makefile":       kafkaMicroservice.KafkaMicroserviceMakefileTemplate,
		"kafka-microservice-docker-compose": kafkaMicroservice.KafkaMicroserviceDockerComposeTemplate,
		"kafka-microservice-broker":         kafkaMicroservice.KafkaMicroserviceBrokerTemplate,

		"rabbitmq-microservice-readme":         rabbitmqMicroservice.RabbitmqMicroserviceReadmeTemplate,
		"rabbitmq-microservice-env":            rabbitmqMicroservice.RabbitmqMicroserviceEnvTemplate,
		"rabbitmq-microservice-makefile":       rabbitmqMicroservice.RabbitmqMicroserviceMakefileTemplate,
		"rabbitmq-microservice-docker-compose": rabbitmqMicroservice.RabbitmqMicroserviceDockerComposeTemplate,
		"rabbitmq-microservice-broker":         rabbitmqMicroservice.RabbitmqMicroserviceBrokerTemplate,
	}
}
