package controllers

import (
	"github.com/revel/revel"
	"class-wechat/app/models"
	"strconv"
)

type NoticeCTL struct {
	*revel.Controller
}

func (c NoticeCTL) Notice() revel.Result {
	return c.Render()
}

func (c NoticeCTL) List() revel.Result {
	search:=c.Params.Get("searchInput")
	page,_:=strconv.Atoi(c.Params.Get("page"))
	list:=models.GetNoticeList(page,10,search)
	return c.RenderJSON(list)
}