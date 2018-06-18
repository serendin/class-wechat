package controllers

import (
	"github.com/revel/revel"
	"class-wechat/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Home() revel.Result {
	if c.Session["username"]==""{
		return c.Redirect("/app/homeForm")
	}
	user:=models.GetUser(c.Session["username"])
	return c.Render(user)
}

func (c App) HomeForm() revel.Result {
	return c.Render()
}

func (c App) BindInfo() revel.Result {
	resp:=models.Resp{Code:0}
	name:=c.Params.Form.Get("name")
	number:=c.Params.Form.Get("number")
	idcard:=c.Params.Form.Get("idcard")
	email:=c.Params.Form.Get("email")
	if name==""||number==""||idcard==""||email==""{
		resp.Msg="输入不完整"
		return c.RenderJSON(resp)
	}
	u:=models.User{Name:name,StuId:number,Email:email,Idcard:idcard}
	if u.CheckUser(){
		c.Session["username"]=u.StuId
		resp.Code=1
		return c.RenderJSON(resp)
	}else{
		resp.Msg="信息匹配错误"
		return c.RenderJSON(resp)
	}
}

func (c App) UnBind() revel.Result {
	c.Session["username"]=""
	return c.Redirect("/app/index")
}
