package main

import (
	"fmt"
	"go-astro/cmd/app"
	"go-astro/internal/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
)

func main() {
	config, err := config.Load(".")
	if err != nil {
		panic(err.Error())
	}

	app := app.New(*config)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Run(); err != nil {
			log.Error(fmt.Sprintf("Error starting server: %s", err.Error()))
		}
	}()
	log.Info(
		fmt.Sprintf(
			"%s is listening on port `%v`",
			config.AppName,
			config.Port,
		),
	)

	<-done

	// some degree of graceful shutdown
	if err := app.Kill(); err != nil {
		log.Error("Failed to kill server: %s", err.Error())
	} else {
		log.Info("Going offline now, catch you later...")
	}
}
