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
	if c.Session["username"]==""{
		return c.Redirect("/app/homeForm")
	}
	return c.Render()
}

func (c QuestionCTL) List() revel.Result {
	search:=c.Params.Get("searchInput")
	page,_:=strconv.Atoi(c.Params.Get("page"))
	list:=models.GetQuestionList(page,10,search,"")
	return c.RenderJSON(list)
}

func (c QuestionCTL) MyList() revel.Result {
	search:=c.Params.Get("searchInput")
	page,_:=strconv.Atoi(c.Params.Get("page"))
	list:=models.GetQuestionList(page,10,search,c.Session["username"])
	return c.RenderJSON(list)
}

func (c QuestionCTL) Ask() revel.Result {
	lessons:=models.GetLessonList(c.Session["username"])
	return c.Render(lessons)
}


func (c QuestionCTL) AskForm() revel.Result {
	resp:=models.Resp{Code:0}
	title:=c.Params.Form.Get("title")
	ques:=c.Params.Form.Get("question")
	lectureId,err:=strconv.Atoi(c.Params.Form.Get("lectureId"))
	if title==""||err!=nil||lectureId==0{
		resp.Msg="输入有误"
		return c.RenderJSON(resp)
	}
	question:=models.Question{Title:title,Question:ques,LectureId:lectureId,
		StuNo:c.Session["username"]}
	err=question.Save()
	if err!=nil{
		resp.Msg="数据库错误"
	}else{
		resp.Code=1
	}
	return c.RenderJSON(resp)
}