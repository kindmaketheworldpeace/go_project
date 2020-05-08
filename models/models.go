package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int
	Email      string
	Password   string
	Admin      bool
	Name       string
	Image      string
	Created_at string
}
type Comment struct {
	Id         int
	Blog_id    string
	User_id    string
	User_name  string
	User_image string
	Content    string
	Created_at string
}

type Blog struct {
	Id         int
	User_id    string
	User_name  string
	User_image string
	Name       string
	Summary    string
	Created_at string
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/go?charset=utf8")
	orm.RegisterModel(new(User), new(Comment), new(Blog))
}
