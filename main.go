package main

import (
	_ "calendar/routers"

	"github.com/astaxie/beego"
	"calendar/models"
	"encoding/gob"
)

func initSession() {
	//beego的session序列号是用gob的方式，因此需要将注册models.User
	gob.Register(models.User{})

	//https://beego.me/docs/mvc/controller/session.md
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName ="SESSION_USER_KEY"
	//beego.BConfig.WebConfig.Session.SessionProvider = "file"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//beego.SetStaticPath("/","views")
	initSession()
	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowAllOrigins:  true,
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	AllowCredentials: true,
	//}))
	beego.BConfig.WebConfig.StaticDir["/statics"] = "statics"
	models.InitDB()
	beego.Run()
}
