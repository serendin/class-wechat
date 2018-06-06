package controllers

import (
	"github.com/revel/revel"
	"class-wechat/app/models"
)

type GradeCTL struct {
	*revel.Controller
}

func (c GradeCTL) Grade() revel.Result {
	return c.Render()
}


func (c GradeCTL) List() revel.Result {
	term:=c.Params.Get("term")
	stuNo:="20090001"
	list:=models.GetGradeList(term,stuNo)
	return c.RenderJSON(list)
}