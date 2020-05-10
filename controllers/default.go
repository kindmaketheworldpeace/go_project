package controllers

import (
	"encoding/json"
	"fmt"
	"go_project/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}
type UserController struct {
	beego.Controller
}
type RegisterController struct {
	beego.Controller
}
type CreateBlogController struct {
	beego.Controller
}
type GetBlogController struct {
	beego.Controller
}
type GetSingleBlogController struct {
	beego.Controller
}
type CreateCommentController struct {
	beego.Controller
}
type GetCommentController struct {
	beego.Controller
}
type LoginController struct {
	beego.Controller
}
type Register struct {
	Name      string
	Email     string
	Password1 string
	password2 string
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *UserController) Get() {
	var users []*models.User
	o := orm.NewOrm()
	_, err := o.QueryTable("user").All(&users)
	fmt.Println(err)
	ret := make(map[string]interface{})
	if err == nil {
		ret["result"] = true
		ret["users"] = &users
		c.Data["json"] = &ret

	} else {
		ret["result"] = false
		ret["message"] = "查询错误！"
	}
	c.ServeJSON()
}

type Blog struct {
	User_Id    string
	User_Name  string
	User_Image string
	Name       string
	Summary    string
	Content    string
}
type Comment struct {
	Blog_id    string
	User_Id    string
	User_Name  string
	User_Image string
	Content    string
}

func (this *CreateBlogController) Post() {

	ob := Blog{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err == nil {
		o := orm.NewOrm()
		var blog models.Blog
		ret := make(map[string]interface{})
		blog.User_id = ob.User_Id
		blog.User_name = ob.User_Name
		blog.User_image = ob.User_Image
		blog.Name = ob.Name
		blog.Summary = ob.Summary
		blog.Content = ob.Content
		blog.Created_at = time.Now().Format("2006-01-02 15:04:05")
		_, err := o.Insert(&blog)
		if err == nil {
			fmt.Println(err)
		}
		ret["result"] = true
		this.Data["json"] = &ret
		this.ServeJSON()

	}

}
func (this *GetBlogController) Get() {
	var blogs []*models.Blog
	o := orm.NewOrm()
	_, err := o.QueryTable("blog").All(&blogs)
	fmt.Println(err)
	ret := make(map[string]interface{})
	if err == nil {
		ret["result"] = true
		ret["blogs"] = &blogs
		this.Data["json"] = &ret

	} else {
		ret["result"] = false
		ret["message"] = "查询错误！"
	}
	this.ServeJSON()

}
func (this *GetSingleBlogController) Get() {
	var comments []*models.Comment
	id := this.GetString("id")
	o := orm.NewOrm()
	_, err := o.QueryTable("comment").Filter("blog_id", id).All(&comments)
	fmt.Println(err)
	blog_id_int, _ := strconv.Atoi(id)
	blog := models.Blog{Id: blog_id_int}
	read_err := o.Read(&blog)
	ret := make(map[string]interface{})
	if read_err == nil {
		ret["result"] = true
		ret["blog"] = &blog
		ret["comments"] = &comments
		this.Data["json"] = &ret

	} else {
		ret["result"] = false
		ret["message"] = "查询错误！"
	}
	this.ServeJSON()

}

func (this *RegisterController) Post() {

	ob := Register{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err == nil {
		o := orm.NewOrm()
		var user models.User
		exist := o.QueryTable("user").Filter("email", ob.Email).Exist()
		ret := make(map[string]interface{})
		if !exist {
			user.Name = ob.Name
			user.Password = ob.Password1
			user.Email = ob.Email
			user.Admin = true
			user.Created_at = time.Now().Format("2006-01-02 15:04:05")
			_, err := o.Insert(&user)
			if err == nil {
				fmt.Println(err)
			}
			ret["result"] = true

		} else {
			ret["result"] = false
			ret["message"] = "失败,已存在相同账户！"

		}
		this.Data["json"] = &ret
		this.ServeJSON()

	}

}
func (this *CreateCommentController) Post() {

	ob := Comment{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err == nil {
		o := orm.NewOrm()
		var comment models.Comment
		ret := make(map[string]interface{})
		comment.User_id = ob.User_Id
		comment.User_name = ob.User_Name
		comment.User_image = ob.User_Image
		comment.Blog_id = ob.Blog_id
		comment.Content = ob.Content
		comment.Created_at = time.Now().Format("2006-01-02 15:04:05")
		_, err := o.Insert(&comment)
		if err == nil {
			fmt.Println(err)
		}
		ret["result"] = true
		this.Data["json"] = &ret
		this.ServeJSON()

	}

}
func (this *GetCommentController) Get() {
	user_id := this.GetString("user_id")
	var comment []*models.Comment
	o := orm.NewOrm()
	_, err := o.QueryTable("comment").Filter("user_id", user_id).All(&comment)
	fmt.Println(err)
	ret := make(map[string]interface{})
	if err == nil {
		ret["result"] = true
		ret["comments"] = &comment
		this.Data["json"] = &ret

	} else {
		ret["result"] = false
		ret["message"] = "查询错误！"
	}
	this.ServeJSON()
}

type Login struct {
	Email  string
	Passwd string
}

func (this *LoginController) Post() {

	login := Login{}
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &login); err == nil {
		o := orm.NewOrm()
		var user models.User

		readerr := o.QueryTable("user").Filter("email", login.Email).One(&user)

		ret := make(map[string]interface{})
		if readerr == orm.ErrNoRows {
			ret["result"] = false
			ret["message"] = "账户不存在!"
		} else {
			if user.Password != login.Passwd {
				ret["result"] = false
				ret["message"] = "密码错误!"

			} else {
				ret["result"] = true
				ret["user"] = &user

			}

		}
		this.Data["json"] = &ret
		this.ServeJSON()
	}
}
