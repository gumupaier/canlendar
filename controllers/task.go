package controllers

import (
	"calendar/models"
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
	"reflect"
)

type TaskController struct {
	BaseController
}

func server_error_task(u *TaskController, code int, data map[string]interface{}) {
	var ret *Response
	var ret1 *ErrorResponse
	if code == 0 {
		ret = &Response{Code: code, Data: data}
		u.Data["json"] = ret
	} else {
		ret1 = &ErrorResponse{Code: code, Msg: data}
		u.Data["json"] = ret1
	}
	u.ServeJSON()
}

type createTask struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// @Title 创建任务
// @Description 创建任务
// @Param	body		body 	createTask	true		"body for user content"
// @Success 200 {object} models.Task
// @Failure 403 body is empty
// @router / [post]
func (u *TaskController) Post() {
	if !u.IsLogin {
		beego.Error("not_has_permission")
		ret := make(map[string]interface{})
		ret["info"] = "not_has_permission"
		server_error_task(u, 400, ret)
	}
	var query createTask
	json.Unmarshal(u.Ctx.Input.RequestBody, &query)

	beego.Info(query)
	title := query.Title
	detail := query.Detail
	fmt.Println(title)
	fmt.Println(detail)
	var task models.Task
	task.Title = title
	task.Detail = detail
	task.CreatePerson = &u.User

	ret, err := models.AddTask(task)
	if err != nil {
		ret := make(map[string]interface{})
		ret["info"] = "add_task_error"
		server_error_task(u, 400, ret)
	}
	ret1 := Struct2Map(*ret)
	server_error_task(u, 0, ret1)
}

// @Title 获取某一任务信息
// @Description 获取某一任务信息
// @Param	id		path 	int	 true	"the task id"
// @Success 200 {object} models.Task
// @Failure 403 body is empty
// @router /:id [get]
func (u *TaskController) GetTask() {
	//if !u.IsLogin {
	//	beego.Error("not_has_permission")
	//	ret := make(map[string]interface{})
	//	ret["info"] = "not_has_permission"
	//	server_error_task(u, 400, ret)
	//}
	id, err := u.GetInt(":id")
	if err != nil {
		fmt.Println(err)
		ret := make(map[string]interface{})
		ret["info"] = "the_params_error"
		server_error_task(u, 400, ret)
	}
	ret, err1 := models.GetTaskById(id)
	if err1 != nil {
		ret := make(map[string]interface{})
		ret["info"] = "get_task_error"
		server_error_task(u, 400, ret)
	}
	//fmt.Println((*ret.CreatePerson).Id)

	result := Struct2Map(*ret)
	server_error_task(u, 0, result)
}

// @Title 更新某一任务（单个）
// @Description 更新任务(代表id的单个任务)
// @Param	id		query 	int	 true	"the task id"
// @Param	title		query 	string	 true	"the task title"
// @Param	detail		query 	string	 true	"the task detail"
// @Success 200 {object} models.Task
// @Failure 403 body is empty
// @router /update/ [put]
func (u *TaskController) Update() {
	if !u.IsLogin {
		beego.Error("not_has_permission")
		ret := make(map[string]interface{})
		ret["info"] = "not_has_permission"
		server_error_task(u, 400, ret)
	}
	id, err := u.GetInt("id")
	if err != nil {

		beego.Error(err)
		ret := make(map[string]interface{})
		ret["info"] = "the_params_error"
		server_error_task(u, 400, ret)
	}
	title := u.GetString("title")
	detail := u.GetString("detail")

	ret, err1 := models.UpdateTask(id, title, detail)
	if err1 != nil {
		ret := make(map[string]interface{})
		ret["info"] = "update_task_error"
		server_error_task(u, 400, ret)
	}
	result := Struct2Map(*ret)
	server_error_task(u, 0, result)
}

const DATE_FORMAT = "2006-01-02"

// @Title 获取当月所有任务
// @Description 获取当月所有任务
// @Param	month		query 	string	 true	"the task title"
// @Param	year		query 	string	 true	"the task title"
// @Success 200 {object} models.Task
// @Failure 403 body is empty
// @router /get_all_task_by_month/ [get]
func (u *TaskController) GetTheMonthAllTask() {
	//u.Prepare()
	if !u.IsLogin {
		beego.Error("not_has_permission")
		ret := make(map[string]interface{})
		ret["info"] = "not_has_permission"
		server_error_task(u, 400, ret)
	}
	//o := orm.NewOrm()
	month, _ := strconv.Atoi(u.GetString("month"))
	year, _ := strconv.Atoi(u.GetString("year"))
	thisMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	//beego.Info(year)
	//beego.Info(month)

	start := thisMonth.AddDate(0, 0, 0).Format(DATE_FORMAT)
	//fmt.Println(start)
	end := thisMonth.AddDate(0, 1, -1).Format(DATE_FORMAT)
	//fmt.Println(end)
	var map_list [] models.TaskMappingResponseResult

	for i := 0; i < 32; i++ {
		start = thisMonth.AddDate(0, 0, i).Format(DATE_FORMAT)
		temp := models.GetTaskMappingByDate(start)
		//if temp.Day==4{
		//
		//	beego.Warning(start)
		//	beego.Warning(temp)
		//}
		//temp.TaskList[0].
		//for _, v := range temp.TaskList {
		//	task_id := v.TaskId
		//	var tasksetting models.TaskSetting
		//	o.QueryTable("task_setting").Filter("task_id", task_id).One(&tasksetting)
		//
		//	//tasksetting.EffectiveDate
		//}
		map_list = append(map_list, temp)
		if start == end {
			break
		}
	}
	u.Data["json"] = &BaseResonse{
		0, map_list,
	}
	u.ServeJSON()
}
func taskIsValid(date string, repeat_time string,effective_date string) bool {
	if repeat_time=="day"{

	}else if repeat_time=="week"{

	}else if repeat_time=="year"{

	}else if repeat_time=="month"{

	}else if repeat_time=="no"{

	}
	return true
}

// @Title 获取当日所有任务
// @Description 获取当月所有任务（end）
// @Param	date		query 	string	 true	"the task title"
// @Success 200 {object} models.Task
// @Failure 403 body is empty
// @router /get_all_task_by_Day/ [get]
func (u *TaskController) GetTheDayTask() {
	if !u.IsLogin {
		beego.Error("not_has_permission")
		ret := make(map[string]interface{})
		ret["info"] = "not_has_permission"
		server_error_task(u, 400, ret)
	}
	date := u.GetString("date")
	result := models.GetTaskMappingByDate(date)
	u.Data["json"] = result
	u.ServeJSON()
}

type editTask struct {
	Title            string `json:"title"`
	EffectiveDate    string `json:"effective_date"`
	Detail           string `json:"detail"`
	RepeatTime       string `json:"repeat_time"`
	AssistPersonList [] int `json:"assist_person_list"`
	RepeatCount      int    `json:"repeat_count"`
	TaskId           int    `json:"task_id"`
}

// @Title 编辑任务
// @Description 编辑任务 (end)
// @Param	body		body 	editTask	true		"The ip for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /edit/ [post]
func (u *TaskController) EditTask() {
	if !u.IsLogin {
		beego.Error("not_has_permission")
		ret := make(map[string]interface{})
		ret["info"] = "not_has_permission"
		server_error_task(u, 400, ret)
	}
	var query editTask
	json.Unmarshal(u.Ctx.Input.RequestBody, &query)
	if len(query.AssistPersonList)==0{
		query.AssistPersonList=append(query.AssistPersonList,u.User.Id)
	}
	o := orm.NewOrm()
	//添加事务
	err := o.Begin()
	//编辑任务分为以下几步
	//Task修改  title detail 其他不变
	//setting修改   repeat_time  ReatCount EffectiveDate Year Month Day WeekDay Enabled
	//mapping添加多条  删除之前该任务对应的mapping  Task AssisPerson TaskSetting

	//1.修改Task
	var data_task models.Task
	task_id := query.TaskId
	data_task.Id = task_id

	err = o.Read(&data_task)
	if err != nil {
		u.Data["json"] = &ErrorResponse{400, "未找到该任务，请检查当前传入参数"}
	}
	data_task.Title = query.Title
	data_task.Detail = query.Detail
	var nn int64
	nn, err = o.Update(&data_task)
	beego.Info("dasdasdasdasd", nn)
	//2.修改setting
	var task_setting models.TaskSetting
	task_setting.Task = &data_task
	err = o.Read(&task_setting, "task_id")
	if err != nil {
		u.Data["json"] = &ErrorResponse{400, "未找到该任务设置，请检查当前传入参数"}
	}
	task_setting.RepeatCount = query.RepeatCount
	task_setting.RepeatTime = query.RepeatTime
	task_setting.EffectiveDate = query.EffectiveDate
	date_li := strings.Split(query.EffectiveDate, "-")
	year, _ := strconv.Atoi(date_li[0])
	month, _ := strconv.Atoi(date_li[1])
	day, _ := strconv.Atoi(date_li[2])
	task_setting.Year = year
	task_setting.Month = month
	task_setting.Day = day
	effect_date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	weekday := int(effect_date.Weekday())
	task_setting.WeekDay = weekday
	_, err = o.Update(&task_setting)

	//3.修改mapping

	//删除该id对应的所有task_mappping
	num, err := o.QueryTable("task_mapping").Filter("task_id", task_id).Delete()
	beego.Info("Affected Num: %s, %s", num, err)
	//根据传来的协助人信息列表插入新的mapping
	for _, v := range query.AssistPersonList {
		var tmp_user models.User
		var tmp_task_mapping models.TaskMapping //新建临时map对象，该对象用于插入
		tmp_user.Id = v
		//o.ReadForUpdate()
		o.Read(&tmp_user)
		tmp_task_mapping.Task = &data_task
		//o.ReadOrCreate(&task_setting,"id")
		beego.Warning(task_setting)
		tmp_task_mapping.TaskSetting = &task_setting
		tmp_task_mapping.AssistPerson = &tmp_user
		_, err = o.Insert(&tmp_task_mapping)
		if err != nil {
			beego.Info("dasdasdasasdasdas", err)
		}
	}
	//事务提交
	err = o.Commit()
	if err != nil {
		//beego.Info("dasdasdasasdasdas", err)
		//提交错误进行事务回滚
		o.Rollback()

		u.Data["json"] = &BaseResonse{401, "更新任务失败，请稍后重试"}
		u.ServeJSON()
		return
	}
	u.Data["json"] = &BaseResonse{0, "update_success"}
	u.ServeJSON()
}

type addTask struct {
	Title            string `json:"title"`
	EffectiveDate    string `json:"effective_date"`
	Detail           string `json:"detail"`
	RepeatTime       string `json:"repeat_time"`
	AssistPersonList [] int `json:"assist_person_list"`
	RepeatCount      int    `json:"repeat_count"`
}

// @Title 添加任务
// @Description 添加任务 (end)
// @Param	body		body 	addTask	true		"The ip for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /add/ [post]
func (u *TaskController) AddTask() {
	//u.Prepare()
	if !u.IsLogin {
		beego.Error("not_has_permission")
		ret := make(map[string]interface{})
		ret["info"] = "not_has_permission"
		server_error_task(u, 400, ret)
	}
	var query editTask
	json.Unmarshal(u.Ctx.Input.RequestBody, &query)
	if len(query.AssistPersonList)==0{
		query.AssistPersonList=append(query.AssistPersonList,u.User.Id)
	}
	o := orm.NewOrm()
	//添加事务
	err := o.Begin()
	//添加任务传入参数与编辑任务一致
	//Task添加  title detail create_person Create_time
	//setting修改   repeat_time  ReatCount EffectiveDate Year Month Day WeekDay Enabled
	//mapping添加多条    Task AssisPerson TaskSetting
	//1.添加Task
	var data_task models.Task

	data_task.Detail = query.Detail
	data_task.Title = query.Title
	data_task.CreatePerson = &u.User
	beego.Info("u.User", u.User)
	data_task.Create_time = int(time.Now().Unix())
	o.Insert(&data_task)
	//2.setting  正对于新建的任务，其setting也是直接添加
	var task_setting models.TaskSetting
	task_setting.Task = &data_task
	task_setting.EffectiveDate = query.EffectiveDate
	date_li := strings.Split(query.EffectiveDate, "-")
	year, _ := strconv.Atoi(date_li[0])
	month, _ := strconv.Atoi(date_li[1])
	day, _ := strconv.Atoi(date_li[2])
	effect_date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	weekday := int(effect_date.Weekday())
	task_setting.EffectiveDate = query.EffectiveDate

	task_setting.Year = year
	task_setting.Month = month
	task_setting.Day = day

	task_setting.WeekDay = weekday
	task_setting.RepeatCount = query.RepeatCount
	task_setting.RepeatTime = query.RepeatTime
	task_setting.Enabled = true
	//beego.Info("333")
	o.Insert(&task_setting)
	//3.修改mapping,其对应mapping也是直接加入
	//根据传来的协助人信息列表插入新的mapping
	beego.Info(len(query.AssistPersonList))
	beego.Info(reflect.TypeOf(query.AssistPersonList))
	for _, v := range query.AssistPersonList {
		var tmp_user models.User
		var tmp_task_mapping models.TaskMapping //新建临时map对象，该对象用于插入
		tmp_user.Id = v
		o.Read(&tmp_user)
		tmp_task_mapping.Task = &data_task
		tmp_task_mapping.AssistPerson = &tmp_user
		tmp_task_mapping.TaskSetting = &task_setting
		//count, err1 :=
		o.Insert(&tmp_task_mapping)
		//beego.Info("插入", count, "条")
		//beego.Info(err1)
		//beego.Info("插入")
	}
	if len(query.AssistPersonList) == 0 {
		//var tmp_user models.User
		var tmp_task_mapping models.TaskMapping //新建临时map对象，该对象用于插入
		//o.Read(&tmp_user)
		tmp_task_mapping.Task = &data_task
		tmp_task_mapping.AssistPerson = &u.User
		tmp_task_mapping.TaskSetting = &task_setting
		//count, err1 :=
		o.Insert(&tmp_task_mapping)
	}
	//beego.Info("346")
	//事务提交
	err = o.Commit()
	//beego.Info("349")
	//beego.Info("err", err)
	if err != nil {
		//提交错误进行事务回滚
		o.Rollback()
		u.Data["json"] = &ErrorResponse{401, "插入数据失败，请稍后重试"}
		return
	}
	u.Data["json"] = &BaseResonse{0, "add_task_success"}
	u.ServeJSON()
}

type sortTaskStruct struct {
	UserName string `json:"user_name"`
	Month    string `json:"month"`
	Year     string `json:"year"`
	Query    string `json:"query"`
}
type sortTaskStructRes struct {
	Day              int    `json:"day"`
	TaskName         string `json:"task_name"`
	AssistPersonList []int  `json:"assist_person_list"`
}

func judge_is_contain(data [] int, query int) bool {
	for _, v := range data {
		if v == query {
			return true
		}
	}
	return false
}

// @Title 获取查询任务
// @Description 根据条件查询当月所有任务
// @Param	month		query 	string	 true	"the task month"
// @Param	year		query 	string	 true	"the task year"
// @Param	username		query 	string	 false	"the task username"
// @Param	query		query 	string	 false	"the task query"
// @Success 200 {object} models.Task
// @Failure 403 body is empty
// @router /get_query_task/ [get]
func (u *TaskController) GetQueryTaskList() {
	if !u.IsLogin {
		beego.Error("not_has_permission")
		ret := make(map[string]interface{})
		ret["info"] = "not_has_permission"
		server_error_task(u, 400, ret)
	}
	o := orm.NewOrm()
	month, _ := strconv.Atoi(u.GetString("month"))
	year, _ := strconv.Atoi(u.GetString("year"))
	username := u.GetString("username")
	query := u.GetString("query")
	thisMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	beego.Info(year)
	beego.Info(month)

	start := thisMonth.AddDate(0, 0, 0).Format(DATE_FORMAT)
	//fmt.Println(start)
	end := thisMonth.AddDate(0, 1, -1).Format(DATE_FORMAT)
	fmt.Println(end)
	var map_list [] models.TaskMappingResponseResult
	var result_list [] interface{}
	for i := 0; i < 32; i++ {
		start = thisMonth.AddDate(0, 0, i).Format(DATE_FORMAT)
		temp := models.GetTaskMappingByDate(start)
		map_list = append(map_list, temp)
		if start == end {
			break
		}
	}

	if username != "" {
		var query_user models.User
		query_user.Username = username
		o.Read(&query_user, "username")
		if query != "" {
			//	两者均不为空
			for _, v := range map_list {
				for _, v1 := range v.TaskList {
					//if v1.TaskTitle
					var tmp_res sortTaskStructRes
					if (strings.Contains(v1.TaskTitle, query) || strings.Contains(v1.TaskDetail, query)) && (v1.CreatePersonName == username || judge_is_contain(v1.AssistPersonList, query_user.Id)) {
						tmp_res.AssistPersonList = v1.AssistPersonList
						tmp_res.Day = v.Day
						tmp_res.TaskName = v1.TaskTitle
						result_list = append(result_list, tmp_res)
					}
				}
			}


		} else if query == "" {
			//	username不为空
			for _, v := range map_list {
				for _, v1 := range v.TaskList {
					var tmp_res sortTaskStructRes
					if (v1.CreatePersonName == username || judge_is_contain(v1.AssistPersonList, query_user.Id)) {
						tmp_res.AssistPersonList = v1.AssistPersonList
						tmp_res.Day = v.Day
						tmp_res.TaskName = v1.TaskTitle
						result_list = append(result_list, tmp_res)
					}
				}
			}

		}
	} else {
		if query != "" {
			for _, v := range map_list {
				for _, v1 := range v.TaskList {
					var tmp_res sortTaskStructRes
					if (strings.Contains(v1.TaskTitle, query) || strings.Contains(v1.TaskDetail, query)) {
						tmp_res.AssistPersonList = v1.AssistPersonList
						tmp_res.Day = v.Day
						tmp_res.TaskName = v1.TaskTitle
						result_list = append(result_list, tmp_res)
					}
				}
			}
		} else if query == "" {
			for _, v := range map_list {
				for _, v1 := range v.TaskList {
					var tmp_res sortTaskStructRes
					tmp_res.AssistPersonList = v1.AssistPersonList
					tmp_res.Day = v.Day
					tmp_res.TaskName = v1.TaskTitle
					result_list = append(result_list, tmp_res)
				}
			}
		}
	}
	//if len(result_list)==0{
	//	var tmp_res sortTaskStructRes
	//	result_list=append(result_list,tmp_res)
	//}
	u.Data["json"] = &BaseResonse{
		0, result_list,
	}
	u.ServeJSON()
}

type deleteTaskStruct struct {
	TaskId int `json:"task_id"`
}

// @Title 删除任务
// @Description 删除任务 (end)
// @Param	body		body 	deleteTaskStruct	true		"The ip for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /delete/ [post]
func (u *TaskController) DeleteTask() {
	if !u.IsLogin {
		beego.Error("not_has_permission")
		ret := make(map[string]interface{})
		ret["info"] = "not_has_permission"
		server_error_task(u, 400, ret)
	}
	var query deleteTaskStruct
	json.Unmarshal(u.Ctx.Input.RequestBody, &query)
	task_id := query.TaskId
	beego.Info(query)
	o := orm.NewOrm()
	//添加事务
	err := o.Begin()
	//var delete_task models.Task
	//var delete_task_setting models.TaskSetting
	//var delete_task_mapping models.TaskMapping
	o.QueryTable("task").Filter("id", task_id).Delete()
	o.QueryTable("task_setting").Filter("task_id", task_id).Delete()
	o.QueryTable("task_mapping").Filter("task_id", task_id).Delete()
	o.Commit()
	if err != nil {
		o.Rollback()
	}
	u.Data["json"] = &BaseResonse{0, "delete_success"}
	u.ServeJSON()
}

type delete_or_edit_task struct {
	IsDelete      bool   `json:"is_delete"`
	IsFinish      bool   `json:"is_finish"`
	EffectiveDate string `json:"effective_date"`
	TaskId        int    `json:"task_id"`
}

// @Title 完成或者删除单个任务
// @Description 完成或者删除单个任务 (end)
// @Param	body		body 	deleteTaskStruct	true		"The ip for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /delete_or_update/ [post]
func (u *TaskController) UpdateOrDeleteTask() {
	if !u.IsLogin {
		beego.Error("not_has_permission")
		ret := make(map[string]interface{})
		ret["info"] = "not_has_permission"
		server_error_task(u, 400, ret)
	}
	var query delete_or_edit_task
	var task_mapping_setting models.TaskMappingSetting
	json.Unmarshal(u.Ctx.Input.RequestBody, &query)
	beego.Warning(query)
	task_id := query.TaskId
	is_delete := query.IsDelete
	is_finish := query.IsFinish
	date := query.EffectiveDate
	beego.Info(query)
	o := orm.NewOrm()
	task_mapping_setting.Date = date
	task_mapping_setting.TaskId = task_id
	err := o.Read(&task_mapping_setting, "date", "task_id")
	if err == nil {
		task_mapping_setting.IsFinish = is_finish
		task_mapping_setting.IsDelete = is_delete
		o.Update(&task_mapping_setting)
	} else if err == orm.ErrNoRows {
		task_mapping_setting.TaskId = task_id
		task_mapping_setting.Date = date
		task_mapping_setting.IsDelete = is_delete
		task_mapping_setting.IsFinish = is_finish
		o.Insert(&task_mapping_setting)
	}
	u.Data["json"] = &BaseResonse{0, "delete_or_update_success"}
	u.ServeJSON()
}
