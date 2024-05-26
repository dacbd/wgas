package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/dacbd/wgas/server"
)

func main() {
	log.Info().Msg("gaa spin up")
	ctx := context.Background()
	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	srv := server.NewServer()
	server := http.Server{
		Addr:    ":4040",
		Handler: srv,
	}
	go func() {
		log.Info().Msg(fmt.Sprintf("listening: %s", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Msg(fmt.Sprintf("error listening and serving: %s", err))
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Error().Msg(fmt.Sprintf("error shutting down http server: %s", err))
		}
	}()
	wg.Wait()
	return nil
}
