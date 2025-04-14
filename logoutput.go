package logoutput

import (
	"os"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"fmt"
)

func main() {
	lf, err := os.OpenFile(
		"logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to open log file")
	}
	multiWriters := zerolog.MultiLevelWriter(os.Stdout, lf)
	log.Logger = zerolog.New(multiWriters).With().Timestamp().Logger()
	log.Info().Msg("Hello World!")
	fmt.Print("Di sini")
}