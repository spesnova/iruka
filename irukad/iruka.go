package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/spesnova/iruka/irukad/controllers"
	"github.com/spesnova/iruka/registry"
	"github.com/spesnova/iruka/scheduler"
)

func main() {
	// Registry
	machines := "http://172.17.8.101:4001"
	reg := registry.NewRegistry(machines, registry.DefaultKeyPrefix)

	// Scheduler
	url := "http://172.17.8.101:4002"
	sch := scheduler.NewScheduler(url)

	// Render
	ren := render.New()

	// Controllers
	appController := controllers.NewAppController(reg, ren)
	containerController := controllers.NewContainerController(reg, ren, sch)

	// Router
	rou := mux.NewRouter()
	v1rou := rou.PathPrefix("/api/v1-alpha").Subrouter()

	// App Resource
	v1rou.Path("/apps").Methods("POST").HandlerFunc(appController.Create)
	v1rou.Path("/apps/{idOrName}").Methods("DELETE").HandlerFunc(appController.Delete)
	v1rou.Path("/apps/{idOrName}").Methods("GET").HandlerFunc(appController.Info)
	v1rou.Path("/apps").Methods("GET").HandlerFunc(appController.List)
	v1rou.Path("/apps/{idOrName}").Methods("PATCH").HandlerFunc(appController.Update)

	v1subrou := v1rou.PathPrefix("/apps/{appIdOrName}").Subrouter()

	// Container Resource
	v1subrou.Path("/containers").Methods("POST").HandlerFunc(containerController.Create)
	v1subrou.Path("/containers/{idOrName}").Methods("DELETE").HandlerFunc(containerController.Delete)
	v1subrou.Path("/containers/{idOrName}").Methods("GET").HandlerFunc(containerController.Info)
	v1subrou.Path("/containers").Methods("GET").HandlerFunc(containerController.List)
	v1subrou.Path("/containers/{idOrName}").Methods("PATCH").HandlerFunc(containerController.Update)

	// Middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	n.UseHandler(rou)

	go containerController.UpdateStates()

	n.Run(":8080")
}
