package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/spesnova/iruka/registry"
	"github.com/spesnova/iruka/scheduler"
	"github.com/spesnova/iruka/schema"
)

const (
	interval = 5
)

type ContainerController struct {
	reg *registry.Registry
	*render.Render
	sch *scheduler.Scheduler
}

func NewContainerController(reg *registry.Registry, ren *render.Render, sch *scheduler.Scheduler) ContainerController {
	return ContainerController{reg, ren, sch}
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

	container, err := c.reg.CreateContainer(appIdOrName, opts)
	if err != nil {
		// TODO (spesnova): if the reqeust is invalid, server should returns 400 instead of 500
		//c.JSON(rw, http.StatusBadRequest, "error")

		// TODO (spesnova): response better error
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	err = c.sch.CreateContainer(container)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(rw, http.StatusCreated, container)
}

func (c *ContainerController) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appIdOrName := vars["appIdOrName"]
	idOrName := vars["idOrName"]

	container, err := c.reg.ContainerFilteredByApp(appIdOrName, idOrName)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = c.reg.DestroyContainer(container.ID.String())
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	err = c.sch.DestroyContainer(container)
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

	container, err := c.reg.ContainerFilteredByApp(appIdOrName, idOrName)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(rw, http.StatusOK, container)
}

func (c *ContainerController) List(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appIdOrName := vars["appIdOrName"]

	containers, err := c.reg.ContainersFilteredByApp(appIdOrName)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	if containers == nil {
		c.JSON(rw, http.StatusOK, []schema.Container{})
		return
	}

	c.JSON(rw, http.StatusOK, containers)
}

func (c *ContainerController) Update(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	appIdOrName := vars["appIdOrName"]
	idOrName := vars["idOrName"]

	var opts schema.ContainerUpdateOpts
	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		// TODO (spesnova): response better error
		c.JSON(rw, http.StatusInternalServerError, "error")
		return
	}

	container, err := c.reg.ContainerFilteredByApp(appIdOrName, idOrName)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	container, err = c.reg.UpdateContainer(idOrName, opts)
	if err != nil {
		c.JSON(rw, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(rw, http.StatusAccepted, container)
}

func (c *ContainerController) UpdateStates() {
	for {
		containers, err := c.sch.Containers()
		if err != nil {
			fmt.Println(err.Error())
		}

		for _, container := range containers {
			_, err := c.reg.UpdateContainerState(container.Name, container.State)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		time.Sleep(interval * time.Second)
	}
}
