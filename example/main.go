package main

import (
	"fmt"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	l "github.com/jesseokeya/go-httplogger"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("src").HTTPBox()))
	http.ListenAndServe(port("8000"), l.Golog(r))
}

func port(p string) string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = p
	}
	fmt.Printf("server running on port *%s \n", port)
	return ":" + port
}
