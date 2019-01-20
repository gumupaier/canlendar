package controllers

import (
	"github.com/astaxie/beego"
	"calendar/models"
	"reflect"
	"strings"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	IsLogin bool        //标识 用户是否登陆
	User    models.User //登陆的用户
}


type Response struct {
	Code int `json:"code"`
	Data map[string]interface{} `json:"data"`
}
type BaseResonse struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
} 
type ErrorResponse struct {
	Code int `json:"code"`
	Msg  interface{} `json:"msg"`
}


func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		//前端要求要转为小写。。。。。
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}




func (ctx *BaseController) Prepare()  {
	// 验证用户是否登陆，判断session中是否存在用户，存在就已经登陆，不存在就没有登陆。
	//sess:=ctx.StartSession()
	ctx.IsLogin = false
	tu:=ctx.GetSession(SESSION_USER_KEY)
	beego.Info("SESSION_USER_KEY Prepare",tu)
	//beego.Info("tu:=sess.Get(SESSION_USER_KEY)",tu)
	//beego.Info("sess.SessionID()",sess.SessionID())
	//tu := ctx.GetSession()
	if tu != nil {
		if u, ok := tu.(models.User); ok {
			ctx.User = u
			ctx.Data["User"] = u
			ctx.IsLogin = true
		}
	}
	ctx.Data["IsLogin"] = ctx.IsLogin
}
