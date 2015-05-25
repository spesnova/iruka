package controllers

import (
	"fmt"
	"net/http"
)

type AppController struct {
}

func (a *AppController) Create(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello")
}

func (a *AppController) Delete(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello")
}

func (a *AppController) Info(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello")
}

func (a *AppController) List(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello")
}

func (a *AppController) Update(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello")
}

func NewAppController() AppController {
	return AppController{}
}
