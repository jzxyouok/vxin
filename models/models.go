package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	connect := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@/" + beego.AppConfig.String("mysqldb") + "?charset=" + beego.AppConfig.String("mysqlcharset")
	// connect := "root:123456@/weichat.vxinapi.cn?charset=utf8"
	// orm.RegisterModelWithPrefix("vxin_", new(User))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", connect)
	// orm.RegisterModel(new(User))
	orm.RegisterModelWithPrefix(beego.AppConfig.String("mysqlprefix"), new(User))
}
