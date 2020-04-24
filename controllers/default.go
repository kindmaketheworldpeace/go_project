package controllers

import (
	"fmt"
	"myproject/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}
type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}
func (c *UserController) Get() {
	o := orm.NewOrm()
	user := models.User{Id: 1}
	err := o.Read(&user)
	c.Data["json"] = user
	c.ServeJSON()
}
