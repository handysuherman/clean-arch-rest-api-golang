package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/docs" //nolint
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/app"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type flags struct {
	appEnvironment string
	configFilePath string
}

// @title Studi Kasus PT. XYZ Golang Developer
// @version 1.0.1
// @description API untuk Studi Kasus PT. XYZ Golang Developer

// @contact.name Handysuherman
// @contact.url https://github.com/handysuherman
// @contact.email lireya95@gmail.com

// @host 0.0.0.0:50050
// @BasePath /api/v1/
// @schemes http

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name x-api-key
// @description API key required to access the API
func main() {
	cmd, err := exportedFlags()
	if err != nil {
		fmt.Println(err)
		return
	}
	logs := logger.NewLogger()
	if cmd.appEnvironment == "develop" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	cfg, err := config.New(cmd.configFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	_app := app.New(logs, cfg)
	logs.Fatal(_app.Run())
}

func exportedFlags() (*flags, error) {
	appConfigFilePath := flag.String("config-file", "./config.yml", "App configuration file path expected extension should be yaml")
	appEnvironment := flag.String("env", "develop", "App environment")

	flag.Parse()

	fileName := filepath.Base(*appConfigFilePath)

	if filepath.Ext(fileName) != ".yaml" {
		return nil, errors.New("configuration should be with .yaml ext")
	}

	return &flags{
		appEnvironment: *appEnvironment,
		configFilePath: *appConfigFilePath,
	}, nil
}
