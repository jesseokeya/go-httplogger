package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(port("3000"), Logger(r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To A Test Server")
}

func port(p string) string {
	port := os.Getenv("PORT")
	if len(port) > 0 {
		p = port
	}
	return ":" + port
}
