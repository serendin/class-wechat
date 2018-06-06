package controllers

import "github.com/revel/revel"

type GradeCTL struct {
	*revel.Controller
}

func (c GradeCTL) Grade() revel.Result {
	return c.Render()
}
