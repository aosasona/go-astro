package main

import (
	"fmt"
	"go-astro/cmd/app"
	"go-astro/configs"
	"go-astro/db"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
)

func main() {
	config, err := configs.LoadEnv(".")
	if err != nil {
		panic(err)
	}

	db, err := db.Connect(db.DatabaseConfig{
		Host:     config.DBHost,
		User:     config.DBUser,
		Password: config.DBPassword,
		DB:       config.DBName,
		Port:     config.DBPort,
	})
	if err != nil {
		panic(err)
	}

	app := app.New(*config, db)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// some degree of graceful shutdown
	go func() {
		_ = <-done
		log.Info("Going offline now, catch you later...")
		_ = app.Kill()
	}()

	log.Info(
		fmt.Sprintf(
			"%s is listening on port `%v`",
			config.AppName,
			config.Port,
		),
	)

	if err := app.Run(); err != nil {
		log.Error(fmt.Sprintf("Error starting server: %s", err.Error()))
	}
}
