package routers

import (
	"AlbumHost/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, `get:Login`)
	beego.Router("/logout", &controllers.UserController{}, `*:Logout`)
	beego.Router("/login", &controllers.UserController{}, `post:PostLogin`)
	beego.Router("/register", &controllers.UserController{}, `get:Register`)
	beego.Router("/register", &controllers.UserController{}, `post:PostRegister`)
}
