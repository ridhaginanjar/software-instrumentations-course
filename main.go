package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"flag"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		oscall := <-ch
		log.Warn().Msgf("system call:%+v", oscall)
		cancel()
	}()

	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	// start: set up any of your logger configuration here if necessary
	logLevel := flag.String("log-level", "info", "Set log level: [debug, info, warn, error, fatal]")

	flag.Parse()

	switch *logLevel {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "warn":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	lf, err := os.OpenFile(
		"logs/app.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666,
	)

	if err != nil {
		log.Fatal().Err(err).Msg("unable to open log file")
	}

	multiWriters := zerolog.MultiLevelWriter(os.Stdout, lf)
	log.Logger = zerolog.New(multiWriters).With().Timestamp().Logger()
	log.Info().Msg("Logging berhasil diaktifkan.")

	// end: set up any of your logger configuration here

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to listen and serve http server")
		}
	}()
	<-ctx.Done()

	if err := server.Shutdown(context.Background()); err != nil {
		log.Error().Err(err).Msg("failed to shutdown http server gracefully")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log := log.With().
		Str("request_id", uuid.New().String()).
		Logger()
	ctx := log.WithContext(r.Context())
	name := r.URL.Query().Get("name")
	res, err := greeting(ctx, name)
	log.Debug().Ctx(ctx).
		Str("func", "handler").
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Str("host", r.Host).
		Str("query", name).
		Str("greeting", res).
		Msg("request received")
	if err != nil {
		log.Error().Ctx(ctx).Msg("Terjadi error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(res))
}

func greeting(ctx context.Context, name string) (string, error) {
	log.Ctx(ctx).Info().
		Str("func", "greeting").
		Msg("greeting function has been called.")

	if len(name) < 5 {
		return fmt.Sprintf("Hello %s! Your name is to short\n", name), nil
	}
	return fmt.Sprintf("Hi %s", name), nil
}
