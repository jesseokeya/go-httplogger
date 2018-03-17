package httplogger

import (
	"net/http"
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("httplogger")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} ▶ %{message}`,
)

func init() {
	errors := logging.NewLogBackend(os.Stderr, "", 0)
	messages := logging.NewLogBackend(os.Stderr, "", 0)

	messagesFormatter := logging.NewBackendFormatter(messages, format)

	backend1Leveled := logging.AddModuleLevel(errors)
	backend1Leveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backend1Leveled, messagesFormatter)
}

// Golog logs requests
func Golog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware(r)
		h.ServeHTTP(w, r)
	})
}

func middleware(r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Debugf("%s %s %s ", r.Proto, r.Method, r.URL)
	case http.MethodPut:
		log.Criticalf("%s %s %s ", r.Proto, r.Method, r.URL)
	case http.MethodHead:
		log.Noticef("%s %s %s ", r.Proto, r.Method, r.URL)
	case http.MethodPost:
		log.Debugf("%s %s %s ", r.Proto, r.Method, r.URL)
	case http.MethodDelete:
		log.Warningf("%s %s %s ", r.Proto, r.Method, r.URL)
	default:
		log.Errorf("%s %s %s ", r.Proto, r.Method, r.URL)
	}
}
