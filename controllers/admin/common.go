package admin

import (
	// "fmt"
	"github.com/astaxie/beego"
)

var Namespace *beego.Namespace

func init() {
	beego.SetStaticPath("/admin_static", "views/admin/static")
	Namespace = beego.NewNamespace("/admin").
		AutoRouter(&AdminController{}).
		AutoRouter(&PublicController{})
}

type CommonController struct {
	beego.Controller
	Uid int64
}

func (this *CommonController) Prepare() {
	this.Uid = this.Checklogin()
}

// 检查用户是否登录，并返回当前登录用户的id
func (this *CommonController) Checklogin() int64 {
	uid := this.GetSession("uid")
	if uid == nil {
		this.Redirect("/admin/public/login", 302)
	}
	uids, _ := uid.(int64)
	return uids
}
