package main

import (
	"context"
	"github.com/patyukin/db/internal/config"
	"github.com/patyukin/db/internal/db"
	"github.com/patyukin/db/internal/dbconn"
	"github.com/patyukin/db/internal/usecase"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("unable to load config: %v", err)
	}

	dbConn, err := dbconn.New(ctx, cfg)
	if err != nil {
		log.Fatal().Msgf("failed connecting to db: %v", err)
	}

	repo := db.New(dbConn)
	uc := usecase.New(repo)
}
