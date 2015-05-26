package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	if opts.Name == "" {
		// TODO (spesnova): response better error
		c.JSON(rw, http.StatusBadRequest, "error")
		return
	}

	app, err := c.CreateApp(opts)
	if err != nil {
		// TODO (spesnova): response better error
		c.JSON(rw, http.StatusInternalServerError, "error")
		return
	}

	c.JSON(rw, http.StatusCreated, app)
}

func (c *AppController) Delete(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello")
}

func (c *AppController) Info(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello")
}

func (c *AppController) List(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello")
}

func (c *AppController) Update(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello")
}
