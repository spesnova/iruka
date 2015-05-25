package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "hello")
	})

	// Middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	n.UseHandler(r)

	n.Run(":8080")
}
