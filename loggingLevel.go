package main

import (
	"flag"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    logLevel := flag.String("log-level", "info", "set log level: [debug, info, warn, error, fatal]")
    
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

	log.Debug().Msg("Ini debug")
	log.Warn().Msg("Ini warning")
	log.Error().Msg("Ini Error")
	log.Fatal().Msg("Ini Fatal")
	log.Info().Msg("Ini Info")
}