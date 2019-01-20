package models
//
//import (
//	"github.com/astaxie/beego/orm"
//	"errors"
//	"github.com/satori/go.uuid"
//)
//
//type Task struct {
//	Id           int
//	Date         string //日期
//	Title        string //标题
//	Detail       string //说明
//	Repeat       string //重复周期 日 周 月 季度 年
//	RepeatTime   string //重复次数
//	CreatePerson * User  `orm:"null;rel(one);on_delete(set_null)"`
//	//创建人
//	AssistPerson *User  `orm:"rel(fk)"` //协助人
//	Enabled      bool   //是否生效
//	Unique_id    string
//}
//
//func init() {
//	orm.RegisterModel(new(Task))
//}
//
//func AddTask(t Task) (*Task, error) {
//	//添加任务
//	o := orm.NewOrm()
//	task := new(Task)
//	task.Date = t.Date
//	task.Title = t.Title
//	task.Detail = t.Detail
//	task.Repeat = t.Repeat
//	task.RepeatTime = t.RepeatTime
//	task.CreatePerson = t.CreatePerson
//	task.AssistPerson = t.AssistPerson
//	task.Enabled = t.Enabled
//
//	uid,err:=uuid.NewV4()
//	if err!=nil{
//		return nil,errors.New("general_uuid_error")
//	}
//	task.Unique_id=uid.String()
//	_, err1 := o.Insert(task)
//	if err1 != nil {
//		return nil, errors.New("insert_error")
//	}
//	return &t, err1
//}
//
//func GetTaskById(id int) (*Task, error) {
//	//根据id查询任务
//	o := orm.NewOrm()
//	task := Task{Id: id}
//	err := o.Read(&task)
//	if err != nil {
//		return nil, err
//	}
//	return &task, err
//}
//
//func GetTaskList(query map[string]string) ([]*Task, error) {
//	//根据参数查询任务
//	o := orm.NewOrm()
//	//task:=Task{Id:id}
//	var task []*Task
//	//err:=o.Read(&task)
//
//	q := o.QueryTable("user")
//	for k, v := range query {
//		if v != "" {
//			q.Filter(k, v)
//		}
//	}
//
//	_, err := q.All(&task)
//
//	if err != nil {
//		return nil, err
//	}
//	return task, err
//}
//
//func UpdateTask(id int, task *Task) (*Task, error) {
//	//更新任务
//	o := orm.NewOrm()
//	t := Task{Id: id}
//	//o.Using("default")
//	if o.Read(&t) == nil {
//		if _, err := o.Update(&t); err != nil {
//			return nil, err
//		}
//	}
//	return &t, nil
//}
//
//func DeleteTask(id int) (int, error) {
//
//	o := orm.NewOrm()
//	ret, err := o.Delete(&Task{Id: id})
//
//	return int(ret), err
//}
