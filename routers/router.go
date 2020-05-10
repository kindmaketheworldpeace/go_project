package routers

import (
	"go_project/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/get_user", &controllers.UserController{})
	beego.Router("/api/register", &controllers.RegisterController{})
	beego.Router("/api/login", &controllers.LoginController{})
	beego.Router("/api/create_blog", &controllers.CreateBlogController{})
	beego.Router("/api/get_blogs", &controllers.GetBlogController{})
	beego.Router("/api/get_blog/", &controllers.GetSingleBlogController{})
	beego.Router("/api/create_comment/", &controllers.CreateCommentController{})
	beego.Router("/api/get_comments/", &controllers.GetCommentController{})
}
