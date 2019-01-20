package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["calendar/controllers:ObjectController"] = append(beego.GlobalControllerRouter["calendar/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:ObjectController"] = append(beego.GlobalControllerRouter["calendar/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:ObjectController"] = append(beego.GlobalControllerRouter["calendar/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:ObjectController"] = append(beego.GlobalControllerRouter["calendar/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:ObjectController"] = append(beego.GlobalControllerRouter["calendar/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:TaskController"] = append(beego.GlobalControllerRouter["calendar/controllers:TaskController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:TaskController"] = append(beego.GlobalControllerRouter["calendar/controllers:TaskController"],
        beego.ControllerComments{
            Method: "GetTask",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:TaskController"] = append(beego.GlobalControllerRouter["calendar/controllers:TaskController"],
        beego.ControllerComments{
            Method: "AddTask",
            Router: `/add/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:TaskController"] = append(beego.GlobalControllerRouter["calendar/controllers:TaskController"],
        beego.ControllerComments{
            Method: "DeleteTask",
            Router: `/delete/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:TaskController"] = append(beego.GlobalControllerRouter["calendar/controllers:TaskController"],
        beego.ControllerComments{
            Method: "EditTask",
            Router: `/edit/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:TaskController"] = append(beego.GlobalControllerRouter["calendar/controllers:TaskController"],
        beego.ControllerComments{
            Method: "GetTheDayTask",
            Router: `/get_all_task_by_Day/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:TaskController"] = append(beego.GlobalControllerRouter["calendar/controllers:TaskController"],
        beego.ControllerComments{
            Method: "GetTheMonthAllTask",
            Router: `/get_all_task_by_month/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:TaskController"] = append(beego.GlobalControllerRouter["calendar/controllers:TaskController"],
        beego.ControllerComments{
            Method: "GetQueryTaskList",
            Router: `/get_query_task/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:TaskController"] = append(beego.GlobalControllerRouter["calendar/controllers:TaskController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:UserController"] = append(beego.GlobalControllerRouter["calendar/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:UserController"] = append(beego.GlobalControllerRouter["calendar/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:UserController"] = append(beego.GlobalControllerRouter["calendar/controllers:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:UserController"] = append(beego.GlobalControllerRouter["calendar/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:UserController"] = append(beego.GlobalControllerRouter["calendar/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetIp",
            Router: `/get_ip`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:UserController"] = append(beego.GlobalControllerRouter["calendar/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserInfo",
            Router: `/get_user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:UserController"] = append(beego.GlobalControllerRouter["calendar/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["calendar/controllers:UserController"] = append(beego.GlobalControllerRouter["calendar/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
