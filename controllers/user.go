package controllers

import (
	"calendar/models"
	"encoding/json"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego"
	"strings"
	"fmt"
	"github.com/astaxie/beego/orm"
)

// Operations about Users
type UserController struct {
	BaseController
}

//func ()server_error_login()

func server_error(u *UserController, code int, data map[string]interface{}) {
	var ret *Response
	var ret1 *ErrorResponse
	if code == 0 {
		ret = &Response{Code: code, Data: data}
		u.Data["json"] = ret
	} else {
		ret1 = &ErrorResponse{Code: code, Msg: data["info"]}
		u.Data["json"] = ret1
	}
	u.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid, err := models.AddUser(user)
	if err != nil {
		u.CustomAbort(400, "add_user_error")
	}
	u.Data["json"] = map[string]int{"id": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = &BaseResonse{Code:0,Data:users}
	u.ServeJSON()
}

// @Title 根据id获取用户信息
// @Description get user by uid
// @Param	id		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @Failure 400 not_has_permission
// @Failure 401 param_error
// @router /:id [get]
//func (u *UserController) Get() {
//	if !u.IsLogin {
//		beego.Error("not_has_permission")
//		u.CustomAbort(400, "not_has_permission")
//	}
//	id, err := u.GetInt(":id")
//	if err != nil {
//		beego.Error("param_error")
//		u.CustomAbort(401, "param_error")
//	}
//	if id != 0 {
//		user := models.GetUserById(id)
//
//		u.Data["json"] = user
//	}
//	u.ServeJSON()
//}

// @Title 更新用户信息
// @Description update the user
// @Param	id		path 	int	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (u *UserController) Update() {
	if !u.IsLogin {
		u.CustomAbort(400, "not_has_permission")
	}
	var user *models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	id, err := u.GetInt(":id")
	if err != nil {
		u.CustomAbort(401, "param_error")
	}
	models.UpdateUser(id, user)
	u.Data["json"] = user
	u.ServeJSON()
}

//@Title 删除用户
//@Description delete the user
//@Param	id		path 	string	true		"The uid you want to delete"
//@Success 200 {string} delete success!
//@Failure 403 id is empty
//@router /:id [delete]
func (u *UserController) Delete() {
	if !u.IsLogin {
		u.CustomAbort(400, "not_has_permission")
	}
	id, err := u.GetInt(":id")
	if err != nil {
		u.CustomAbort(401, "param_error")
	}
	models.DeleteUser(id)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description 登录 (end)
// @Param	body		body 	models.User	true		"The ip for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {

	//beego.BeeLogger.Debug(username)
	//beego.BeeLogger.Debug()

	var user models.User

	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//ip := u.Ctx.Input.GetData("ip")
	//username := u.Ctx.Input.GetData("username")
	ip := user.Ip
	username := user.Username
	//fmt.Println(user)

	//username := u.GetString("username")
	if user, err := models.Login(ip, username); err == nil {
		sess:=u.StartSession()
		sess.Set(SESSION_USER_KEY,user)
		beego.Info("login SESSION_USER_KEY",sess.Get(SESSION_USER_KEY))
		//u.SetSession(SESSION_USER_KEY, user)
		data := Struct2Map(user)
		server_error(u, 0, data)
		u.Data["json"] = user
	} else {
		fmt.Println(err)
		msg := make(map[string]interface{})
		msg["info"] = "请输入正确的用户名"
		server_error(u, 400, msg)
	}
	ret_res := Struct2Map(user)
	server_error(u, 0, ret_res)
	//u.ServeJSON()
}

// @Title 退出登录
// @Description Logs out current logged in user session(end)
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	ret_data := make(map[string]interface{})
	ret_data["status"] = "logout success"
	ret := &Response{Code: 0, Data: ret_data}

	u.Data["json"] = ret
	u.DelSession(SESSION_USER_KEY)
	u.ServeJSON()
}

// @Title 获取ip
// @Description 获取ip(end)
// @Success 200 {string} "127.0.0.1"
// @router /get_ip [get]
func (u *UserController) GetIp() {
	req := u.Ctx.Request
	addr := req.RemoteAddr
	ip := strings.Split(addr, ":")[0]
	beego.BeeLogger.Debug(ip)
	ip_res := make(map[string]interface{})
	ip_res["ip"] = ip
	server_error(u, 0, ip_res)
}

type UserInfoResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Ip       string `json:"ip"`
}

// @Title 获取用户列表
// @Description 根据姓名获取用户列表(end)
// @Param	username		query 	string	 true	" the param for user_list"
// @Success 200 {string} "127.0.0.1"
// @router /get_user [get]
func (u *UserController) GetUserInfo() {
	o := orm.NewOrm()
	username := u.GetString("username")
	var users []* models.User
	var ret [] UserInfoResponse
	var userInfo UserInfoResponse
	o.QueryTable("user").Filter("username", username).All(&users)
	for _, v := range users {
		userInfo.Id = v.Id
		userInfo.Username = v.Username
		userInfo.Ip = v.Ip
		ret = append(ret, userInfo)
	}
	u.Data["json"] = &BaseResonse{
		0,ret,
	}
	u.ServeJSON()
}


