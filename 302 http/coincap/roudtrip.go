package coincap

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Базовый Middleware для логгирования запросов
type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %sn", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}
