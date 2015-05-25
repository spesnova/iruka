package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/spesnova/iruka/irukad/controllers"
)

func main() {
	r := mux.NewRouter()

	// Controllers
	appController := controllers.NewAppController()

	// App Resource
	r.Path("/apps").Methods("POST").HandlerFunc(appController.Create)
	r.Path("/apps/{idOrName}").Methods("DELETE").HandlerFunc(appController.Delete)
	r.Path("/apps/{idOrName}").Methods("GET").HandlerFunc(appController.Info)
	r.Path("/apps").Methods("GET").HandlerFunc(appController.List)
	r.Path("/apps/{idOrName}").Methods("PATCH").HandlerFunc(appController.Update)

	// Middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	n.UseHandler(r)

	n.Run(":8080")
}
