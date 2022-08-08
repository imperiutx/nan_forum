package main

import (
	"database/sql"
	"github.com/imperiutx/nan_forum/api"
	db "github.com/imperiutx/nan_forum/db/sqlc"
	"github.com/imperiutx/nan_forum/utils"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Println("cannot load config:", err)
		return err
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Println("cannot connect to db:", err)
		return err
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Println("cannot create server:", err)
		return err
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Println("cannot start server:", err)
		return err
	}

	return nil
}
