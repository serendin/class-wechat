package controllers

import (
	"github.com/revel/revel"
	"class-wechat/app/models"
	"strconv"
)

type QuestionCTL struct {
	*revel.Controller
}

func (c QuestionCTL) Question() revel.Result {
	return c.Render()
}

func (c QuestionCTL) List() revel.Result {
	search:=c.Params.Get("searchInput")
	page,_:=strconv.Atoi(c.Params.Get("page"))
	list:=models.GetQuestionList(page,10,search)
	return c.RenderJSON(list)
}