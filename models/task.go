package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
	"strings"
	"strconv"
	"time"
	"github.com/astaxie/beego"
)

const DATE_FORMAT = "2006-01-02"

func init() {
	orm.RegisterModel(new(Task))
	orm.RegisterModel(new(TaskSetting))
	orm.RegisterModel(new(TaskMapping))
	orm.RegisterModel(new(TaskMappingSetting))
}

type Task struct {
	Id           int            `json:"id"`
	Title        string         `json:"title"`                                                //标题
	Detail       string         `json:"detail"`                                               //说明
	CreatePerson *User          `orm:"null;rel(fk);on_delete(set_null)" json:"create_person"` //创建人
	Create_time  int            `json:"create_time"`                                          //创建时间
	TaskSetting  *TaskSetting   `orm:"null;rel(one);on_delete(set_null)"json:"task_setting"`
	TaskMappings []*TaskMapping `orm:"reverse(many)"json:"task_mappings"`
	//IsFinished  bool `json:"is_finished"`
}

type TaskSetting struct {
	Id            int    `json:"id"`
	RepeatTime    string `json:"repeat_time"`    //重复周期 日 周 月 季度 年  事实上后来我觉得这个名字换一个比较好
	RepeatCount   int    `json:"repeat_count"`   //重复次数
	EffectiveDate string `json:"effective_date"` //生效时间   为某一日  如2017-08-09
	//将effectiveDate年月日拆分，便于查询，相比于构造复杂的查询来说，这想法完美
	Year     int   `json:"year"`
	Month    int   `json:"month"`
	Day      int   `json:"day"`
	WeekDay  int   `json:"week_day"`
	Task     *Task `orm:"rel(one)"` //任务
	Enabled  bool  `json:"enabled"` //是否生效
	IsFinsh  bool  `json:"is_finsh"`
	IsDelete bool  `json:"is_delete"`
	//TaskMappings []*TaskMapping `orm:"reverse(many)"`
}

type TaskMapping struct {
	Id           int          `json:"id"`
	Task         *Task        `orm:"rel(fk)" json:"task"`
	AssistPerson *User        `orm:"rel(fk)" json:"assist_person"`
	TaskSetting  *TaskSetting `orm:"rel(fk)" json:"task_setting"`
	//TaskMappingSetting *TaskMappingSetting `orm:"null;rel(fk)" json:"task_setting"`
}

type TaskMappingSetting struct {
	Id int `json:"id"`
	//TaskMappings [] * TaskMapping `json:"task_mappings"`
	//TaskMappingId int `json:"task_mapping_id"`
	Date     string
	IsDelete bool `json:"is_delete"`
	IsFinish bool `json:"is_finish"`
	TaskId   int  `json:"task_id"`
}

/*************************任务*****************************/
//添加任务
func AddTask(t Task) (*Task, error) {
	o := orm.NewOrm()
	task := new(Task)
	task.Title = t.Title
	task.Detail = t.Detail
	task.CreatePerson = t.CreatePerson
	task.Create_time = t.Create_time
	_, err := o.Insert(task)
	if err != nil {
		beego.Error(err)
		return nil, errors.New("insert_error")
	}
	return task, err
}

//根据id查询任务
func GetTaskById(id int) (*Task, error) {
	o := orm.NewOrm()
	task := Task{Id: id}
	err := o.Read(&task)
	if err != nil {
		return nil, err
	}
	return &task, err
}

//根据参数查询任务，返回为任务列表
func GetTaskList(query map[string]string) ([]*Task, error) {
	//根据参数查询任务
	o := orm.NewOrm()
	//task:=Task{Id:id}
	var task []*Task
	//err:=o.Read(&task)

	q := o.QueryTable("user")
	for k, v := range query {
		if v != "" {
			q.Filter(k, v)
		}
	}

	_, err := q.All(&task)

	if err != nil {
		return nil, err
	}
	return task, err
}

//更新任务
func UpdateTask(id int, title, detail string) (*Task, error) {
	o := orm.NewOrm()
	t := Task{Id: id}

	if o.Read(&t) == nil {
		if title != "" {
			t.Title = title
		}
		if detail != "" {
			t.Detail = detail
		}
		if _, err := o.Update(&t); err != nil {
			return nil, err
		}
	}
	return &t, nil
}

//删除任务
func DeleteTask(id int) (int, error) {
	o := orm.NewOrm()
	ret, err := o.Delete(&Task{Id: id})
	return int(ret), err
}

/*************************任务设置*****************************/
//添加任务设置
func AddTaskSetting(task_id int, setting TaskSetting) (*TaskSetting, error) {
	o := orm.NewOrm()
	task := Task{Id: task_id}
	err := o.Read(&task)
	if err != nil {
		return nil, err
	}
	task_setting := TaskSetting{Task: &task}
	err1 := o.Read(&task_setting)
	date_li := strings.Split(setting.EffectiveDate, "-")
	year, _ := strconv.Atoi(date_li[0])
	month, _ := strconv.Atoi(date_li[1])
	day, _ := strconv.Atoi(date_li[2])

	effect_date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	weekday := int(effect_date.Weekday())
	if err1 == orm.ErrNoRows {
		//	如果没有设置，则为插入操作
		task_setting = TaskSetting{}
		task_setting.RepeatTime = setting.RepeatTime
		task_setting.RepeatCount = setting.RepeatCount
		task_setting.Task = &task
		task_setting.EffectiveDate = setting.EffectiveDate
		task_setting.Year = year
		task_setting.Month = month
		task_setting.Day = day
		task_setting.WeekDay = weekday
		task_setting.Enabled = setting.Enabled
		_, err = o.Insert(task_setting)
	} else if err1 == nil {
		//如果没有设置，则为更新操作
		task_setting.RepeatTime = setting.RepeatTime
		task_setting.RepeatCount = setting.RepeatCount

		if setting.EffectiveDate != "" {
			task_setting.EffectiveDate = setting.EffectiveDate
			task_setting.Year = year
			task_setting.Month = month
			task_setting.Day = day
			task_setting.WeekDay = weekday
		}

		_, err = o.Update(&task_setting)
	}

	if err == nil {
		return &task_setting, nil
	} else {
		return nil, err
	}
}

//根据任务设置id查看任务设置
func GetTaskSettingById(id int) (*TaskSetting, error) {
	o := orm.NewOrm()
	task_setting := TaskSetting{Id: id}
	err := o.Read(&task_setting)
	if err != nil {
		return nil, err
	}
	return &TaskSetting{}, err
}

//删除任务设置
func DeleteTaskSetting(id int) (int, error) {
	o := orm.NewOrm()
	ret, err := o.Delete(&TaskSetting{Id: id})
	return int(ret), err
}

/*************************任务mapping日期*****************************/
//添加任务协助人，这里传入的是一个user列表
func AddTaskAssist(task_id int, user_list []*User) (int, error) {
	o := orm.NewOrm()
	task := Task{Id: task_id}
	err := o.Read(&task)
	if err != nil {
		return 0, errors.New("the_task_not_found")
	}
	task_setting := TaskSetting{Task: &task}
	err = o.Read(&task_setting)
	if err != nil {
		return 0, errors.New("the_task_setting_not_found")
	}
	var task_mapping_list = make([] TaskMapping, 0, len(user_list))
	for i, v := range task_mapping_list {
		v.AssistPerson = user_list[i]
		v.Task = &task
		v.TaskSetting = &task_setting
	}
	err = o.Begin()
	var successNums int64
	successNums, err = o.InsertMulti(len(user_list), task_mapping_list)
	//事务操作，如果没有全部添加成功返回错误,回滚数据库
	if int(successNums) != len(user_list) {
		o.Rollback()
		return 0, errors.New("insert_error")
	}
	return int(successNums), err
}

type TaskMappingResponse struct {
	TaskMappingList [] TaskMapping `json:"task_mapping_list"`
	Date            string         `json:"date"`
}

type TaskMappingRes struct {
	TaskId           int    `json:"task_id"`
	AssistPersonList []int  `json:"assist_person_list"`
	TaskTitle        string `json:"task_title"`
	TaskDetail       string `json:"task_detail"`
	RepeatTime       string `json:"repeat_time"`
	CreatePersonName string `json:"create_person_name"`
	RepeatCount      int    `json:"repeat_count"`
	IsDelete         bool   `json:"is_delete"`
	IsFinish         bool   `json:"is_finish"`
}

type TaskMappingResponseResult struct {
	Day int `json:"day"`
	//Date     string           `json:"date"`
	TaskList []TaskMappingRes `json:"task_list"`
}

//搜索任务mapping
func GetTaskMappingByDate(date string) TaskMappingResponseResult {
	//orm.Debug = true
	//根据时间搜索任务mapping，分为以下几种情况
	//1、设置日期，不循环
	//2、每天循环，必显示
	//3、每周循环  创建当日的weekday与当前时间weekday相同
	//4、每月循环  创建当日的日与今天的日相同  如 7月1日与8月1日，代表两者相同
	//5、每年循环  创建当日的月和日与今天的月和日相同  如 2017年8月1日与2018年8月1日
	//6、自定义天数 今日是否为创建当日时间的整数倍，即取模为0，如设置每隔n天执行，则(curday-create_day)%setting_day==0
	o := orm.NewOrm()
	orm.Debug=true
	//自定义表达式
	cond := orm.NewCondition()
	date_li := strings.Split(date, "-")
	year, _ := strconv.Atoi(date_li[0])
	month, _ := strconv.Atoi(date_li[1])
	day, _ := strconv.Atoi(date_li[2])
	effect_date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	weekday := int(effect_date.Weekday())

	//year, month, day := time.Now().Date()
	//cur_date=time.Date(year, month, day, 0, 0, 0, 0, time.Local).Format(DATE_FORMAT)
	var task_mapping_list [] TaskMapping
	//不循环及日循环
	cond1 := cond.Or("tasksetting__EffectiveDate", date).Or("tasksetting__RepeatTime", "day")

	//周循环
	cond2 := cond.And("tasksetting__weekday", weekday).And("tasksetting__RepeatTime", "week")
	//月循环
	cond3 := cond.And("tasksetting__RepeatTime", "month").And("tasksetting__day", day)
	//年
	cond4 := cond.And("tasksetting__RepeatTime", "year").And("tasksetting__Month", month).And("tasksetting__Day", day)
	cond5 := cond.OrCond(cond1).OrCond(cond2).OrCond(cond3).OrCond(cond4)

	cond_bak := cond.OrCond(cond.And("tasksetting__year__lt", year)).OrCond(cond.And("tasksetting__Year", year).And("tasksetting__Month__lt", month)).OrCond(cond.And("tasksetting__Year", year).And("tasksetting__Month", month).And("tasksetting__Day__lte", day))

	cond10 := cond.AndCond(cond5).AndCond(cond_bak)

	o.QueryTable("task_mapping").SetCond(cond10).OrderBy("-Id").All(&task_mapping_list)
	if len(task_mapping_list) == 0 {
		var task_mapping_response_result TaskMappingResponseResult
		task_mapping_response_result.Day = day
		//beego.Warning("没有查到任何数据")
		return task_mapping_response_result
	} else {
		var task_mapping_response_result TaskMappingResponseResult
		//beego.Warning(task_mapping_list)
		for _, v := range task_mapping_list {
			var task Task
			var user User
			var task_setting TaskSetting
			var create_person User
			task.Id = v.Task.Id
			o.Read(&task)
			user.Id = v.AssistPerson.Id
			o.Read(&user)
			create_person.Id = task.CreatePerson.Id
			o.Read(&create_person)

			task_setting.Id = v.TaskSetting.Id
			o.Read(&task_setting)

			//beego.Info("task", task.Id)
			//beego.Info("CreatePerson", task.CreatePerson)
			//beego.Info("TaskSetting", v.TaskSetting.Id)

			flag := false //设置标志位，判断该任务是否已在返回列表中

			for i, v1 := range task_mapping_response_result.TaskList {

				if v1.TaskId == task.Id {
					//添加到协助人列表
					task_mapping_response_result.TaskList[i].AssistPersonList = append(task_mapping_response_result.TaskList[i].AssistPersonList, user.Id)
					flag = true
				}

			}

			//需组合的结构体格式如下
			//type TaskMappingRes struct {
			//	Task_id          int      `json:"task_id"`
			//	AssistPersonList []string `json:"assist_person_list"`
			//	TaskTitle        string   `json:"task_title"`
			//	TaskDetail       string   `json:"task_detail"`
			//	RepeatTime       string   `json:"repeat_time"`
			//	CreatePersonName string   `json:"create_person_name"`
			//}
			if !flag {
				//如果进来的话说明该任务不在列表中，直接添加便是
				//task_mapping_response_result.TaskList[i].AssistPersonList=append(task_mapping_response_result.TaskList[i].AssistPersonList,v.AssistPerson.Username)
				var task_mapping_res TaskMappingRes
				var task_mapping_setting_obj TaskMappingSetting
				task_mapping_setting_obj.TaskId = task.Id
				task_mapping_setting_obj.Date = date
				err := o.QueryTable("task_mapping_setting").Filter("task_id", task.Id).Filter("date", date).One(&task_mapping_setting_obj)
				//err:=o.Read(&task_mapping_setting_obj,"task_id","date")

				//if task.Id==6{
				//	beego.Warning("这是任务6")
				//	beego.Warning(task.Id)
				//	beego.Warning(date)
				//	beego.Warning(task_mapping_setting_obj.IsDelete)
				//	beego.Warning(task_mapping_setting_obj.IsFinish)
				//}
				if err == nil {
					task_mapping_res.IsDelete = task_mapping_setting_obj.IsDelete
					task_mapping_res.IsFinish = task_mapping_setting_obj.IsFinish
				}
				task_mapping_res.AssistPersonList = append(task_mapping_res.AssistPersonList, user.Id)
				task_mapping_res.CreatePersonName = create_person.Username
				task_mapping_res.TaskId = task.Id
				task_mapping_res.RepeatTime = task_setting.RepeatTime
				task_mapping_res.TaskDetail = task.Detail
				task_mapping_res.TaskTitle = task.Title
				task_mapping_res.RepeatCount = task_setting.RepeatCount
				//将上面组合好的task加入到返回结果列表中
				task_mapping_response_result.Day = day
				if !task_mapping_setting_obj.IsDelete {
					task_mapping_response_result.TaskList = append(task_mapping_response_result.TaskList, task_mapping_res)
				}
				flag = false
			}
		}
		//if date=="2018-12-04"{
		//	beego.Info(task_mapping_response_result)
		//	beego.Info(date)
		//}
		//beego.Info(date)
		return task_mapping_response_result
	}
}

/******************************************************/

//func IsRepeat(date,effective_date string,repeat_time string) bool{
//	time_tmp:="2018-10-10"
//	t1,_:=time.Parse(time_tmp,date)
//	t2,_:=time.Parse(time_tmp,effective_date)
//
//
//	if repeat_time=="no"{
//
//	}else if repeat_time=="day"{
//
//
//	}else if repeat_time=="week"{
//
//	}else if repeat_time=="month"{
//
//	}else if repeat_time=="year"{
//
//	}
//
//
//	return true
//}
