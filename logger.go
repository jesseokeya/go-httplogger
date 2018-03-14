package httpLogger

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

func httpLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware(r)
		h.ServeHTTP(w, r)
	})
}

func middleware(r *http.Request) {
	switch r.Method {
	case "GET":
		log.Debugf("%s %s %s ", r.Proto, r.Method, r.URL)
	case "PUT":
		log.Criticalf("%s %s %s ", r.Proto, r.Method, r.URL)
	case "HEAD":
		log.Noticef("%s %s %s ", r.Proto, r.Method, r.URL)
	case "POST":
		log.Debugf("%s %s %s ", r.Proto, r.Method, r.URL)
	case "DELETE":
		log.Warningf("%s %s %s ", r.Proto, r.Method, r.URL)
	default:
		log.Errorf("%s %s %s ", r.Proto, r.Method, r.URL)
	}
}
