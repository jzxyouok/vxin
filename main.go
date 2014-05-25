package main

import (
	"github.com/astaxie/beego"
	"vxin/controllers"
	"vxin/controllers/admin"
)

func init() {
	beego.SetStaticPath("/ExtJs", "static/plugins/ExtJs")
}

func main() {
	beego.NewNamespace()
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/admin", &admin.AdminController{}, "*:Index;get:Show")
	beego.AutoRouter(&admin.AdminController{})
	beego.AutoRouter(&admin.PublicController{})

	ns := beego.NewNamespace()

	beego.Run()
}
