package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/spesnova/iruka/irukad/controllers"
	"github.com/spesnova/iruka/registry"
)

func main() {
	// Registry
	machines := "http://172.17.8.101:4001"
	reg := registry.NewRegistry(machines, registry.DefaultKeyPrefix)

	// Render
	ren := render.New()

	// Router
	rou := mux.NewRouter()

	// Controllers
	appController := controllers.NewAppController(reg, ren)
	containerController := controllers.NewContainerController(reg, ren)

	// App Resource
	rou.Path("/apps").Methods("POST").HandlerFunc(appController.Create)
	rou.Path("/apps/{idOrName}").Methods("DELETE").HandlerFunc(appController.Delete)
	rou.Path("/apps/{idOrName}").Methods("GET").HandlerFunc(appController.Info)
	rou.Path("/apps").Methods("GET").HandlerFunc(appController.List)
	rou.Path("/apps/{idOrName}").Methods("PATCH").HandlerFunc(appController.Update)

	subr := rou.PathPrefix("/apps/{appIdOrName}").Subrouter()

	// Container Resource
	subr.Path("/containers").Methods("POST").HandlerFunc(containerController.Create)
	//subr.Path("/containers/{idOrName}").Methods("DELETE").HandlerFunc(containerController.Delete)
	//subr.Path("/containers/{idOrName}").Methods("GET").HandlerFunc(containerController.Info)
	//subr.Path("/containers").Methods("GET").HandlerFunc(containerController.List)
	//subr.Path("/containers/{idOrName}").Methods("PATCH").HandlerFunc(containerController.Update)

	// Middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	n.UseHandler(rou)

	n.Run(":8080")
}
