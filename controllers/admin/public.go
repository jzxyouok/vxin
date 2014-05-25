package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"vxin/common"
	"vxin/models"
)

type PublicController struct {
	beego.Controller
}

func (this *PublicController) Logintest() {
	o := orm.NewOrm()
	user_model := models.User{Id: 1}
	// user_model.Id = 4
	// user_model.Id = 2
	// user_model.LastLocation = "few"
	// user_model.LastLoginIp = "ip"
	// user_model.LastLoginTime = 123
	// user_model.Password = "pwd"
	// user_model.Role = 12
	// user_model.Status = 1
	// id, _ := o.Insert(user_model)
	o.Read(&user_model)
	fmt.Println(user_model)
	this.TplNames = "admin/login.tpl"
	// this.Render()
	this.Data["json"] = []interface{}{"name", "吴赐有", "info", map[string]int{"age": 123}}
	this.ServeJson()
	fmt.Println(common.Md5("aaa"))
}
func (this *PublicController) Login() {
	this.TplNames = "admin/login.tpl"
	this.Render()
}

// 用户注册方法
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
			fmt.Println(o.Insert(&user))
			fmt.Println("register ok")
			this.Data["json"] = map[string]interface{}{"res": "ok", "msg": "恭喜您，用户注册成功！"}
		} else {
			fmt.Println("register error")
		}
		this.ServeJson()
	}
}

// 检查用户登录
func (this *PublicController) CheckLogin() {
	result := map[string]interface{}{"res": "no", "redirect": "/admin/admin/index"}
	user := models.User{Username: this.GetString("username")}
	if err := orm.NewOrm().Read(&user, "Username"); err == nil {
		if user.Password == common.Md5(this.GetString("password")) {
			result["res"] = "yes"
		}
	}
	this.Data["json"] = result
	this.ServeJson()
}

func (this *PublicController) CheckAccess() {

}
