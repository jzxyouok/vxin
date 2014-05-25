package admin

type AdminController struct {
	CommonController
}

func (this *AdminController) Index() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "admin/demo.tpl"
	this.Render()
}

func (this *AdminController) Show() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "admin/admin.tpl"
}
