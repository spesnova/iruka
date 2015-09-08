package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/spesnova/iruka/registry"
	"github.com/spesnova/iruka/router"
	"github.com/spesnova/iruka/schema"
)

type DomainController struct {
	reg *registry.Registry
	*render.Render
	rou *router.Router
}

func NewDomainController(reg *registry.Registry, ren *render.Render, rou *router.Router) DomainController {
	return DomainController{reg, ren, rou}
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

	app, err := c.reg.App(appIdentity)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
	}

	routes, err := c.reg.RoutesFilteredByApp(appIdentity)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
	}

	for _, route := range routes {
		// TODO(dtan4):
		//   Remove duplicates (most routes are the same)
		err := c.rou.AddRoute(app.ID.String(), opts.Hostname, route.Location)

		if err != nil {
			c.JSON(rw, http.StatusInternalServerError, err.Error())
		}
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

	app, err := c.reg.App(appIdentity)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
	}

	domain, err := c.reg.DomainFilteredByApp(appIdentity, identity)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
	}

	err = c.rou.RemoveRoute(app.ID.String(), domain.Hostname)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
	}

	_, err = c.reg.DestroyDomain(identity)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
	}

	c.JSON(rw, http.StatusAccepted, domain)
}

func (c *DomainController) Info(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appIdentity := vars["appIdentity"]
	identity := vars["identity"]

	domain, err := c.reg.DomainFilteredByApp(appIdentity, identity)

	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
	}

	c.JSON(rw, http.StatusOK, domain)
}

func (c *DomainController) List(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appIdentity := vars["appIdentity"]

	domains, err := c.reg.DomainsFilteredByApp(appIdentity)

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
