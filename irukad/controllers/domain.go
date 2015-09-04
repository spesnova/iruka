package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/spesnova/iruka/registry"
	"github.com/spesnova/iruka/schema"
)

type DomainController struct {
	reg *registry.Registry
	*render.Render
}

func NewDomainController(reg *registry.Registry, ren *render.Render) DomainController {
	return DomainController{reg, ren}
}

func (c *DomainController) Create(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	appIdentity := vars["appIdentity"]

	var opts schema.DomainCreateOpts
	err := json.NewDecoder(r.Body).Decode(&opts)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
	}

	domain, err := c.reg.CreateDomain(appIdentity, opts)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
	}

	c.JSON(rw, http.StatusCreated, domain)
}

func (c *DomainController) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appIdentity := vars["appIdentity"]
	identity := vars["identity"]

	domain, err := c.reg.DestroyDomain(appIdentity, identity)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
	}

	c.JSON(rw, http.StatusAccepted, domain)
}

func (c *DomainController) Info(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appIdentity := vars["appIdentity"]
	identity := vars["identity"]

	domain, err := c.reg.Domain(appIdentity, identity)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
	}

	c.JSON(rw, http.StatusOK, domain)
}

func (c *DomainController) List(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appIdentity := vars["appIdentity"]

	domains, err := c.reg.Domains(appIdentity)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	if domains == nil {
		c.JSON(rw, http.StatusOK, []schema.Domain{})
		return
	}

	c.JSON(rw, http.StatusOK, domains)
}

func (c *DomainController) Update(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	appIdentity := vars["appIdentity"]
	identity := vars["identity"]

	var opts schema.DomainUpdateOpts
	err := json.NewDecoder(r.Body).Decode(&opts)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	domain, err := c.reg.Domain(appIdentity, identity)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	domain, err = c.reg.UpdateDomain(appIdentity, identity, opts)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(rw, http.StatusAccepted, domain)
}
