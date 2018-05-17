package controllers

import (
	"github.com/revel/revel"
	"class-wechat/app/models"
	"strconv"
)

type WorkCTL struct {
	*revel.Controller
}

func (c WorkCTL) Work() revel.Result {
	return c.Render()
}

func (c WorkCTL) List() revel.Result {
	search:=c.Params.Get("searchInput")
	page,_:=strconv.Atoi(c.Params.Get("page"))
	list:=models.GetWorkList(page,10,search)
	return c.RenderJSON(list)
}