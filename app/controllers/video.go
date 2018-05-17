package controllers

import (
	"github.com/revel/revel"
	"class-wechat/app/models"
	"strconv"
)

type VideoCTL struct {
	*revel.Controller
}

func (c VideoCTL) Video() revel.Result {
	return c.Render()
}

func (c VideoCTL) List() revel.Result {
	name:=c.Params.Get("searchInput")
	page,_:=strconv.Atoi(c.Params.Get("page"))
	list:=models.GetVideoList(page,10,name)
	return c.RenderJSON(list)
}

func (c VideoCTL) Detail() revel.Result {
	id,_:=strconv.Atoi(c.Params.Get("id"))
    result:=models.GetVideoById(id)
	return c.Render(result)
}