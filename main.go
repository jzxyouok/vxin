package main

import (
	"github.com/astaxie/beego"
	"vxin/controllers/admin"
)

func init() {
	beego.SetStaticPath("/ExtJs", "static/plugins/ExtJs")
	beego.SetStaticPath("/static", "static")
}

// main function
func main() {
	// beego.Router("/", &controllers.MainController{})
	// // beego.Router("/admin", &admin.AdminController{}, "*:Index;get:Show")
	// beego.AutoRouter(&admin.AdminController{})
	// beego.AutoRouter(&admin.PublicController{})
	// ns := beego.NewNamespace()
	//
	beego.AddNamespace(admin.Namespace)
	beego.Run()
}
