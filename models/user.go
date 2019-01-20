package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
	"github.com/jinzhu/gorm"
	"errors"
)

var (
	//定义db变量，变量名开头是小写，只能同包下访问
	db gorm.DB
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int  `json:"id"`
	Username string `json:"username"`
	Password string  `json:"password"`
	Ip       string  `json:"ip"`
	//设置反向关系
	Tasks     []*Task `orm:"reverse(many)"`
	TaskMappings     []*TaskMapping `orm:"reverse(many)"`
}

func AddUser(u User) (int, error) {
	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username
	user.Ip = u.Ip
	user.Password = u.Password

	id, err := o.Insert(user)
	return int(id), err
}

func GetUserById(id int) User {
	u := User{Id: id}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		fmt.Println("查无此人")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	fmt.Println(u)
	return u
}

func GetAllUsers() []*User {
	o := orm.NewOrm()
	o.Using("default")
	var users []*User
	q := o.QueryTable("user")
	q.All(&users)
	return users
}

func UpdateUser(id int, user *User) *User {
	o := orm.NewOrm()
	u := User{Id: id}
	//o.Using("default")
	//Username string
	//Password string
	//Ip       string
	if user.Username != "" {
		u.Username = user.Username
	}
	if user.Password != "" {
		u.Password = user.Password
	}
	if user.Ip != "" {
		u.Ip = user.Ip
	}
	if o.Read(&u) == nil {
		if _, err := o.Update(&u); err != nil {
			return &u
		}
	}
	return nil
}

func DeleteUser(id int) {
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&User{Id: id})
}

func Login(ip, name string) (User, error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("ip", ip).One(&user)
	if err == orm.ErrNoRows {
		user.Username=name
		user.Ip=ip
		o.Insert(&user)
		return user, nil
	}
	fmt.Println(user)
	if strings.Compare(name, user.Username) != 0 {
		//fmt.Println(name)
		//fmt.Println(user.Username)
		return user, errors.New("the_password_wrong")
	}
	return user, nil
}
