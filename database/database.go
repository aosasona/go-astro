package database

import (
	"go-astro/configs"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pocketbase/dbx"
)

var Conn *dbx.DB

func init() {
	db, err := dbx.Open("mysql", configs.GlobalConfig.DSN)
	if err != nil {
		panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	if configs.GlobalConfig.AppEnv == configs.DEVELOPMENT {
		db.LogFunc = log.Printf
	}

	Conn = db
}
