package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("httplogger")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

// Redacted password log
func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	// For demo purposes, create two backend for os.Stderr.
	errors := logging.NewLogBackend(os.Stderr, "", 0)
	messages := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	messagesFormatter := logging.NewBackendFormatter(messages, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(errors)
	backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, messagesFormatter)

	log.Debugf("debug %s", Password("secret"))
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit")
	http.ListenAndServe(":8000", Middleware(r))
}

// Middleware handler interface
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println("middleware", r.URL)
		log.Debugf("%s %s", r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}

// HomeHandler handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my test server")
}
