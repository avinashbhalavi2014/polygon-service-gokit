package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"polygon-service-gokit/polygonApi"
	"polygon-service-gokit/util"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

func main() {

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// load configuration settings from config file
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Error().Err(err).Msg("can't read config file")
		os.Exit(1)
	}

	log.Info().Msgf("service %s started", config.ServiceName)
	defer log.Info().Msg("service ended")

	// connect to database
	db, err := sqlx.Connect(config.DBDriver, config.DBConnectionString)
	if err != nil {
		log.Error().Err(err).Str("msg", "could not connect to database").Msg("exit")
		os.Exit(-1)
	} else {
		log.Info().Msg("Postgresql DB connected successfully")
	}

	// run db migration
	if err := runDBMigration(config.MigrationURL, config.DBConnectionString); err != nil {
		log.Error().Err(err).Msg("exit")
		os.Exit(1)
	} else {
		log.Info().Msg("db migrated successfully")
	}

	ctx := context.Background()

	var srv polygonApi.Service
	{
		repo := polygonApi.NewRepo(db)
		srv = polygonApi.NewService(repo, log)
	}

	endpoints := polygonApi.MakeEndpoints(srv)

	handler := polygonApi.NewHttpServer(ctx, endpoints)

	// server
	log.Info().Msgf("ServerListeing on port :%s", config.Port)
	if err := http.ListenAndServe(config.HostIp+":"+config.Port, handler); err != nil {
		log.Fatal().Err(err).Str("service", config.ServiceName).Msgf("Cannot start %s service", config.ServiceName)
	}

}

func runDBMigration(migrationURL string, dbSource string) error {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		return errors.New("cannot create new migration instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migration up: %v", err)
	}

	return nil
}
