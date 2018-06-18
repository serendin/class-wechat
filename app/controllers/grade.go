package controllers

import (
	"github.com/revel/revel"
	"class-wechat/app/models"
)

type GradeCTL struct {
	*revel.Controller
}

func (c GradeCTL) Grade() revel.Result {
	if c.Session["username"]==""{
		return c.Redirect("/app/homeForm")
	}
	return c.Render()
}


func (c GradeCTL) List() revel.Result {
	term:=c.Params.Get("term")
	stuNo:=c.Session["username"]
	list:=models.GetGradeList(term,stuNo)
	return c.RenderJSON(list)
}