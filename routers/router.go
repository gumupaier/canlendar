// @APIVersion 1.0.0
// @Title the_frist_blood
// @Description beego很牛逼
// @Contact henson_wu@foxmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"calendar/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "DELETE", "PUT", "PATCH", "OPTIONS","POST"},
		AllowHeaders: []string{"Origin", "Access-Control-Allow-Origin","Access-Control-Allow-Headers","Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin","Access-Control-Allow-Headers",},
		AllowCredentials: true,
	}))
	beego.Get("/", func(ctx *context.Context){
		ctx.Redirect(302, "/statics/index.html")
	})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/task",
			beego.NSInclude(
				&controllers.TaskController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
