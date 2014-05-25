package admin

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/admin_static", "views/admin/static")
}

type CommonController struct {
	beego.Controller
}
