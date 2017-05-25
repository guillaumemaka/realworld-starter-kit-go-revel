package controllers

import (
	"github.com/revel/revel"
)

type UserController struct {
	ApplicationController
}

func (c UserController) GetUser() revel.Result {
	return c.Todo()
}

func (c UserController) Register() revel.Result {
	return c.Todo()
}

func (c UserController) Login() revel.Result {
	return c.Todo()
}

func (c UserController) UpdateUser() revel.Result {
	return c.Todo()
}
