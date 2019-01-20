package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() error {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	err := orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RunSyncdb("default", false, true)
	return err
}
