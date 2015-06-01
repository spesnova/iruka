package fleet

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	unitsPath = "/units"
)

type Unit struct {
	Name         string        `json:"name,omitempty"`
	Options      []*UnitOption `json:"options,omitempty"`
	DesiredState string        `json:"desiredState,omitempty"`
	CurrentState string        `json:"currentState,omitempty"`
	MachineID    string        `json:"machineID,omitempty"`
}

type UnitOption struct {
	Section string `json:"section,omitempty"`
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
}

type unitResponse struct {
	Units []Unit `json:"units,omitempty"`
}

func (c *Client) Units() ([]Unit, error) {
	var unitsRes unitResponse

	req, err := http.NewRequest("GET", c.URL+basePath+unitsPath, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&unitsRes)
	if err != nil {
		return nil, err
	}

	return unitsRes.Units, nil
}

// createOrUpdateUnit creates or updates an unit
func (c *Client) createOrUpdateUnit(u Unit) error {
	switch u.DesiredState {
	case "inactive":
		break
	case "loaded":
		break
	case "launched":
		break
	default:
		return errors.New("Invalid desired state")
	}

	j, err := json.Marshal(u)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.URL+basePath+unitsPath+"/"+u.Name, bytes.NewReader(j))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 201:
		// Created successfully
		return nil
	case 204:
		// Modified successfully
		return nil
	case 400:
		// Attempting to create/modify an Unit with an invalid entity
		return errors.New("400 Bad Request")
	case 409:
		// Attempting to create an entity without options
		return errors.New("409 Conflict")
	default:
		message := fmt.Sprintf("%d Faild to create/update an unit", res.StatusCode)
		return errors.New(message)
	}

	return nil
}

func (c *Client) deleteUnit(name string) error {
	req, err := http.NewRequest("DELETE", c.URL+basePath+unitsPath+"/"+name, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 204:
		// Deleted successfully
		return nil
	case 404:
		// The indicated Unit does not exist
		return errors.New("404 Not Found")
	default:
		message := fmt.Sprintf("%d Faild to delete an unit", res.StatusCode)
		return errors.New(message)
	}

	return nil
}

func (c *Client) Submit(name string, opts []*UnitOption, targetState string) error {
	unit := Unit{
		Name:         name,
		Options:      opts,
		DesiredState: targetState,
	}

	return c.createOrUpdateUnit(unit)
}

func (c *Client) Load(name string) error {
	unit := Unit{
		Name:         name,
		DesiredState: "loaded",
	}

	return c.createOrUpdateUnit(unit)
}

func (c *Client) Start(name string) error {
	unit := Unit{
		Name:         name,
		DesiredState: "launched",
	}

	return c.createOrUpdateUnit(unit)
}

func (c *Client) Stop(name string) error {
	unit := Unit{
		Name:         name,
		DesiredState: "loaded",
	}

	return c.createOrUpdateUnit(unit)
}

func (c *Client) Unload(name string) error {
	unit := Unit{
		Name:         name,
		DesiredState: "inactive",
	}

	return c.createOrUpdateUnit(unit)
}

func (c *Client) Destroy(name string) error {
	return c.deleteUnit(name)
}
