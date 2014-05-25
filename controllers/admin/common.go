package admin

import (
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
}
