package main

import (
	_ "AlbumHost/models"
	_ "AlbumHost/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./data/test.db")
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "AlbumHost"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"

	beego.Run()
}
