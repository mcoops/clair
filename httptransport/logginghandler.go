package httptransport

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

type httpStatusWriter struct {
	http.ResponseWriter

	StatusCode int
}

// LoggingHandler will log HTTP requests using the pre initialized zerolog.
func LoggingHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log := zerolog.Ctx(r.Context())

		next.ServeHTTP(w, r)

		out := fmt.Sprintf("handled HTTP request: remote_addr=%s method=%s request_uri=%s status=%s elapsed_time_(ms)=%f",
			r.RemoteAddr, r.Method, r.RequestURI, strconv.Itoa(http.StatusOK), float64(time.Since(start).Nanoseconds())*1e-6)

		log.Info().Msg(out)
	}
}
