package admin

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"vxin/common"
	"vxin/models"
)

type PublicController struct {
	beego.Controller
}

// 用户登出
func (this *PublicController) LoginOut() {
	this.DelSession("uid")
	this.Redirect("/admin/public/login", 200)
}

// 用户登录
func (this *PublicController) Login() {
	method := this.Ctx.Input.Method()
	if method == "GET" {
		this.TplNames = "admin/login.tpl"
		this.Render()
	} else {
		result := map[string]interface{}{"res": "no"}
		user := models.User{Username: this.GetString("username")}
		if err := orm.NewOrm().Read(&user, "Username"); err == nil {
			if user.Password == common.Md5(this.GetString("password")) {
				result = map[string]interface{}{"res": "yes", "redirectUrl": "/admin/admin/index"}
				this.SetSession("uid", user.Id)
			}
		}
		this.Data["json"] = result
		this.ServeJson()
	}
}

// 用户注册
func (this *PublicController) Register() {
	method := this.Ctx.Input.Method()
	if method == "GET" {
		this.TplNames = "admin/register.tpl"
		this.Render()

	} else if method == "POST" {
		this.Data["json"] = map[string]interface{}{"res": "err", "msg": "注册失败,用户名已经存在！"}
		o := orm.NewOrm()
		user := models.User{
			Username: this.GetString("username"),
			Password: common.Md5(this.GetString("password")),
			Email:    this.GetString("email"),
		}
		if err := o.Read(&user, "Username"); err != nil {
			userid, _ := o.Insert(&user)
			this.Data["json"] = map[string]interface{}{
				"res":         "ok",
				"msg":         "恭喜您，用户注册成功！",
				"redirectUrl": "/admin/admin/index",
			}
			this.SetSession("uid", userid)
		}
		this.ServeJson()
	}
}
