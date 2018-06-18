package controllers

import (
	"github.com/revel/revel"
	"class-wechat/app/models"
	"strconv"
)

type MaterialCTL struct {
	*revel.Controller
}

func (c MaterialCTL) Material() revel.Result {
	if c.Session["username"]==""{
		return c.Redirect("/app/homeForm")
	}
	return c.Render()
}

func (c MaterialCTL) List() revel.Result {
	search:=c.Params.Get("searchInput")
	page,_:=strconv.Atoi(c.Params.Get("page"))
	list:=models.GetMaterialList(page,10,search)
	return c.RenderJSON(list)
}