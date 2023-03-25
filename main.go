package main

import (
	"fmt"
	"go-astro/cmd/app"
	"go-astro/configs"
	"go-astro/database"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"xorm.io/xorm"
)

var (
	db     *xorm.Engine
	config *configs.Config

	log *zap.Logger
	err error
)

func init() {
	config, err = configs.LoadEnv(".")
	if err != nil {
		panic(err)
	}

	db, err = database.Connect(database.DatabaseConfig{
		Host:     config.DBHost,
		User:     config.DBUser,
		Password: config.DBPassword,
		DB:       config.DBName,
		Port:     config.DBPort,
	})
	if err != nil {
		panic(err)
	}

	// Create new logger instance
	if config.AppEnv == configs.DEVELOPMENT {
		log, err = zap.NewDevelopment()
	} else {
		log, err = zap.NewProduction()
	}

	if err != nil {
		panic("failed to create logger")
	}
}

func main() {
	gastro := app.New(*config, db)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// some degree of graceful shutdown
	go func() {
		_ = <-done
		log.Info("Going offline now, catch you later...")
		_ = gastro.Kill()
	}()

	log.Info(
		fmt.Sprintf(
			"%s is listening on port `%v`",
			config.AppName,
			config.Port,
		),
	)

	if err := gastro.Run(); err != nil {
		log.Error(fmt.Sprintf("Error starting server: %s", err.Error()))
	}
}
