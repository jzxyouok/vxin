package admin

type AdminController struct {
	CommonController
}

func (this *AdminController) Index() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "admin/admin.tpl"
}

func (this *AdminController) Show() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	// this.TplNames = "admin/admin.tpl"
	this.TplNames = "admin/demo.tpl"
}

func (this *AdminController) Login() {
	this.TplNames = "admin/login.tpl"
}
