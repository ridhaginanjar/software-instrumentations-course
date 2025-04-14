package simpleLog 

import (
	//"github.com/rs/zerolog"
	"errors"
	"github.com/rs/zerolog/log"
)

func main() {
	err := errors.New("payment is already expired!")
	log.Fatal().
		Err(err).
		Str("payment_id", "123").
		Str("payment_status", "failed").
		Str("booking_id", "98989").
		Float64("amount", 100.0).
		Msgf("payment failed for booking %s", "98989")
}