package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/spesnova/iruka/registry"
	"github.com/spesnova/iruka/schema"
)

type ConfigVarsController struct {
	*registry.Registry
	*render.Render
}

func NewConfigVarsController(reg *registry.Registry, ren *render.Render) ConfigVarsController {
	return ConfigVarsController{reg, ren}
}

func (c *ConfigVarsController) Create(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	appIdentity := vars["appIdentity"]

	var opts schema.ConfigVars
	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, "error")
		return
	}

	app, err := c.CreateConfigVars(appIdentity, opts)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, "error")
		return
	}

	c.JSON(rw, http.StatusCreated, app)
}
