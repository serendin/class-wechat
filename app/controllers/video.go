package controllers

import (
	"github.com/revel/revel"
)

type VideoCTL struct {
	*revel.Controller
}

func (c VideoCTL) Index() revel.Result {
	return c.Render()
}
