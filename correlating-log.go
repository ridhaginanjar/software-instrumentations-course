package correlatingLog

import (
	"context"
	"fmt"
	"net/http"
	"html"
	"github.com/rs/zerolog/log"
	"github.com/google/uuid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	log := log.With().
		Str("request_id", uuid.New().String()).
		Logger()
	ctx := log.WithContext(r.Context())
	log.Info().Ctx(ctx).
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Str("host", r.Host).
		Str("query", query).
		Msg("request received")
	doFirst(ctx)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func doFirst(ctx context.Context) {
	log := log.Ctx(ctx)
	log.Info().Msg("do first")
	doSecond(ctx)
}

func doSecond(ctx context.Context) {
	log.Ctx(ctx).Info().Msg("do second")
}

func main() {
	http.HandleFunc("/", handler)
	log.Info().Msg("localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Server failed to run!")
	}
}