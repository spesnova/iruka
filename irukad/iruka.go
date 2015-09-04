package main

import (
	"log"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/spesnova/iruka/agent"
	"github.com/spesnova/iruka/irukad/controllers"
	"github.com/spesnova/iruka/registry"
	"github.com/spesnova/iruka/router"
	"github.com/spesnova/iruka/scheduler"
)

func main() {
	// Registry
	reg := registry.NewRegistry(registry.DefaultMachines, registry.DefaultKeyPrefix)

	// Sub-domain router
	rou := router.NewRouter(router.DefaultMachines, router.DefaultKeyPrefix)

	// Scheduler
	url := scheduler.DefaultAPIURL
	sch := scheduler.NewScheduler(url)

	// Agent
	machine := os.Getenv("IRUKA_MACHINE")
	if machine == "" {
		log.Fatal("IRUKA_MACHINE is required, but missing")
	}
	age := agent.NewAgent(agent.DefaultHost, machine, reg)

	// Render
	ren := render.New()

	// Controllers
	appController := controllers.NewAppController(reg, ren, rou)
	containerController := controllers.NewContainerController(reg, ren, sch)
	configVarsController := controllers.NewConfigVarsController(reg, ren)
	domainController := controllers.NewDomainController(reg, ren, rou)

	// Router
	muxRou := mux.NewRouter()
	v1rou := muxRou.PathPrefix("/api/v1-alpha").Subrouter()

	// App Resource
	v1rou.Path("/apps").Methods("POST").HandlerFunc(appController.Create)
	v1rou.Path("/apps/{identity}").Methods("DELETE").HandlerFunc(appController.Delete)
	v1rou.Path("/apps/{identity}").Methods("GET").HandlerFunc(appController.Info)
	v1rou.Path("/apps").Methods("GET").HandlerFunc(appController.List)
	v1rou.Path("/apps/{identity}").Methods("PATCH").HandlerFunc(appController.Update)

	v1subrou := v1rou.PathPrefix("/apps/{appIdentity}").Subrouter()

	// Container Resource
	v1subrou.Path("/containers").Methods("POST").HandlerFunc(containerController.Create)
	v1subrou.Path("/containers/{identity}").Methods("DELETE").HandlerFunc(containerController.Delete)
	v1subrou.Path("/containers/{identity}").Methods("GET").HandlerFunc(containerController.Info)
	v1subrou.Path("/containers").Methods("GET").HandlerFunc(containerController.List)
	v1subrou.Path("/containers/{identity}").Methods("PATCH").HandlerFunc(containerController.Update)

	// Config Vars Resource
	v1subrou.Path("/config-vars").Methods("GET").HandlerFunc(configVarsController.Info)
	v1subrou.Path("/config-vars").Methods("PATCH").HandlerFunc(configVarsController.Update)

	// Domain Resource
	v1subrou.Path("/domains").Methods("POST").HandlerFunc(domainController.Create)
	v1subrou.Path("/domains/{identity}").Methods("DELETE").HandlerFunc(domainController.Delete)
	v1subrou.Path("/domains/{identity}").Methods("GET").HandlerFunc(domainController.Info)
	v1subrou.Path("/domains").Methods("GET").HandlerFunc(domainController.List)

	// Middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	n.UseHandler(muxRou)

	go age.Pulse()
	// Disable retrieving unit state from fleet for now
	//go containerController.UpdateStates()

	n.Run(":8080")
}
