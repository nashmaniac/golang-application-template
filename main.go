package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/nashmaniac/golang-application-template/config"
	"github.com/nashmaniac/golang-application-template/endpoints/app"
	"github.com/nashmaniac/golang-application-template/infrastructure/postgres"
	"github.com/nashmaniac/golang-application-template/usecases"
	"github.com/sherifabdlnaby/configuro"
)

func readConfiguration() (*config.Config, error) {
	wd, _ := os.Getwd()
	configPath := path.Join(wd, "config")
	env, ok := os.LookupEnv("environment")
	if !ok {
		env = "local"
	}
	configFilePath := fmt.Sprintf("%s/config.%s.yaml", configPath, env)
	configOptions := configuro.WithLoadFromConfigFile(configFilePath, true)
	conf, err := configuro.NewConfig(configOptions)
	if err != nil {
		return nil, err
	}
	c := &config.Config{}
	err = conf.Load(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func main() {

	ctx := context.Background()

	// read the configs
	configuration, err := readConfiguration()
	if err != nil {
		panic(err)
	}

	// first initialize the db repository
	dbRepository, err := postgres.NewRepository(
		ctx,
		configuration.Database.Username,
		configuration.Database.Password,
		configuration.Database.Name,
		configuration.Database.Host,
		configuration.Database.Port,
	)
	if err != nil {
		panic(err)
	}

	defer dbRepository.CloseDB(ctx)

	usecases, err := usecases.NewUsecases(
		ctx,
		dbRepository,
		configuration,
	)
	if err != nil {
		panic(err)
	}

	endpoints, err := app.NewEndpoints(usecases, configuration)
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("running server on port %d", configuration.Application.Port))
	endpoints.Server.Run(fmt.Sprintf(":%d", configuration.Application.Port))

}
