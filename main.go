package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/imperiutx/nan_forum/api"
	db "github.com/imperiutx/nan_forum/db/sqlc"
	"github.com/imperiutx/nan_forum/utils"
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

	if err = runDBMigration(config.MigrationURL, config.DBSource); err != nil {
		log.Println("cannot migrate:", err)
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

func runDBMigration(migrationURL string, dbSource string) error {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Println("cannot create new migrate instance:", err)
		return err
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Println("failed to run migrate up:", err)
		return err
	}

	log.Println("db migrated successfully")

	return nil
}
