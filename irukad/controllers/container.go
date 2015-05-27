package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/spesnova/iruka/registry"
	"github.com/spesnova/iruka/schema"
)

type ContainerController struct {
	*registry.Registry
	*render.Render
}

func NewContainerController(reg *registry.Registry, ren *render.Render) ContainerController {
	return ContainerController{reg, ren}
}

func (c *ContainerController) Create(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	appIdOrName := vars["appIdOrName"]

	var opts schema.ContainerCreateOpts
	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	container, err := c.CreateContainer(appIdOrName, opts)
	if err != nil {
		// TODO (spesnova): if the reqeust is invalid, server should returns 400 instead of 500
		//c.JSON(rw, http.StatusBadRequest, "error")

		// TODO (spesnova): response better error
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(rw, http.StatusCreated, container)
}

func (c *ContainerController) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appIdOrName := vars["appIdOrName"]
	idOrName := vars["idOrName"]

	container, err := c.ContainerFilteredByApp(appIdOrName, idOrName)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = c.DestroyContainer(container.ID.String())
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(rw, http.StatusAccepted, container)
}

func (c *ContainerController) Info(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appIdOrName := vars["appIdOrName"]
	idOrName := vars["idOrName"]

	container, err := c.ContainerFilteredByApp(appIdOrName, idOrName)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(rw, http.StatusOK, container)
}

func (c *ContainerController) List(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appIdOrName := vars["appIdOrName"]

	containers, err := c.ContainersFilteredByApp(appIdOrName)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(rw, http.StatusOK, containers)
}
