package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/spesnova/iruka/registry"
	"github.com/spesnova/iruka/schema"
)

type AppController struct {
	*registry.Registry
	*render.Render
}

func NewAppController(reg *registry.Registry, ren *render.Render) AppController {
	return AppController{reg, ren}
}

func (c *AppController) Create(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var opts schema.AppCreateOpts
	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		// TODO (spesnova): response better error
		c.JSON(rw, http.StatusInternalServerError, "error")
		return
	}

	app, err := c.CreateApp(opts)
	if err != nil {
		// TODO (spesnova): if the reqeust is invalid, server should returns 400 instead of 500
		//c.JSON(rw, http.StatusBadRequest, "error")

		// TODO (spesnova): response better error
		c.JSON(rw, http.StatusInternalServerError, "error")
		return
	}

	ropts := schema.RouteCreateOpts{
		Location: "/.*",
		Upstream: app.ID.String(),
	}

	if _, err := c.Registry.CreateRoute(app.ID.String(), ropts); err != nil {
		c.JSON(rw, http.StatusInternalServerError, "error")
		return
	}

	c.JSON(rw, http.StatusCreated, app)
}

func (c *AppController) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	identity := vars["identity"]

	app, err := c.DestroyApp(identity)
	if err != nil {
		// TODO (spesnova): separate 404 and 500 error
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(rw, http.StatusAccepted, app)
}

func (c *AppController) Info(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	identity := vars["identity"]

	app, err := c.App(identity)
	if err != nil {
		// TODO (spesnova): separate 404 and 500 error
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(rw, http.StatusOK, app)
}

func (c *AppController) List(rw http.ResponseWriter, r *http.Request) {
	apps, err := c.Apps()
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, "error")
		return
	}

	if len(apps) == 0 {
		c.JSON(rw, http.StatusOK, []schema.App{})
		return
	}

	c.JSON(rw, http.StatusOK, apps)
}

func (c *AppController) Update(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	identity := vars["identity"]

	var opts schema.AppUpdateOpts
	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		// TODO (spesnova): response better error
		c.JSON(rw, http.StatusInternalServerError, "error")
		return
	}

	app, err := c.UpdateApp(identity, opts)
	if err != nil {
		// TODO (spesnova): if the reqeust is invalid, server should returns 400 instead of 500
		//c.JSON(rw, http.StatusBadRequest, "error")

		// TODO (spesnova): response better error
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(rw, http.StatusAccepted, app)
}
