package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/imperiutx/nan_forum/api"
	db "github.com/imperiutx/nan_forum/db/sqlc"
	"github.com/imperiutx/nan_forum/pb"
	"github.com/imperiutx/nan_forum/gapi"
	"github.com/imperiutx/nan_forum/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	if err = runGrpcServer(config, store); err != nil {
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

func runGrpcServer(config utils.Config, store db.Store) error {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		// log.Fatal().Err(err).Msg("cannot create server")
		return err
	}

	// gprcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	// grpcServer := grpc.NewServer(gprcLogger)
	grpcServer := grpc.NewServer()
	pb.RegisterNanForumServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		// log.Fatal().Err(err).Msg("cannot create listener")
		return err
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())

	// log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		// log.Fatal().Err(err).Msg("cannot start gRPC server")
		return err
	}

	return nil
}

func runRestServer(config utils.Config, store db.Store) error {
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
