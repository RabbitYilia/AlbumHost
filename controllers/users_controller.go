package controllers

import (
	"AlbumHost/models"
	"log"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	username := c.GetSession("Username")
	Respjson := map[string]interface{}{"Server": "AlbumHost Server"}
	if username == nil {
		Respjson["IsLogin"] = false
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "login.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	} else {
		c.Data["IsLogin"] = true
		c.Data["Username"] = username
		Respjson["IsLogin"] = true
		Respjson["Username"] = username
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "index.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	}
}

func (c *UserController) Logout() {
	username := c.GetSession("Username")
	Respjson := map[string]interface{}{"Server": "AlbumHost Server"}
	if username == nil {
		c.Data["IsLogin"] = false
		Respjson["IsLogin"] = false
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "index.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	} else {
		c.DelSession("Username")
		c.Data["IsLogin"] = false
		Respjson["IsLogin"] = false
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "index.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	}
}

func (c *UserController) Register() {
	username := c.GetSession("Username")
	Respjson := map[string]interface{}{"Server": "AlbumHost Server"}
	if username == nil {
		Respjson["IsLogin"] = false
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "register.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	} else {
		c.Data["IsLogin"] = true
		c.Data["Username"] = username
		Respjson["IsLogin"] = true
		Respjson["Username"] = username
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "index.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	}
}

func (c *UserController) PostRegister() {
	username := c.GetSession("Username")
	Respjson := map[string]interface{}{"Server": "AlbumHost Server"}
	if username == nil {
		username := c.GetString("Username") // register.html中传过来的数据
		password := c.GetString("Password")
		log.Printf("[Register][Received Info]username=%s,pwd=%s\n", username, password)

		valid := validation.Validation{}
		valid.Required(username, "Username") // 校验是否为空值
		valid.Required(password, "Password")
		// valid.MaxSize(id, 20, "id")
		switch { // 使用switch方式来判断是否出现错误，如果有错，则打印错误并返回
		case valid.HasErrors():
			log.Printf("[Register][Validation]%s\n", valid.Errors[0].Key+valid.Errors[0].Message)
			c.Data["Error"] = "注册失败：" + valid.Errors[0].Key + valid.Errors[0].Message
			Respjson["Error"] = "注册失败：" + valid.Errors[0].Key + valid.Errors[0].Message
			if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
				c.Data["IsLogin"] = false
				c.TplName = "bad.html"
			} else {
				c.Data["json"] = Respjson
				c.ServeJSON()
			}
			return
		}

		u := models.User{
			Username: username,
		}
		err := u.ReadByName()
		if err == nil {
			log.Printf("[Register][Err]%s\n", "User Already Exists")
			c.Data["IsLogin"] = false
			c.Data["Error"] = "注册失败：User Already Exists"
			Respjson["IsLogin"] = false
			Respjson["Error"] = "注册失败：User Already Exists"
			if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
				c.Data["IsLogin"] = false
				c.TplName = "bad.html"
			} else {
				c.Data["json"] = Respjson
				c.ServeJSON()
			}
			return
		}

		nu := models.User{
			Username: username,
			Password: password,
		}
		err = nu.Create()
		if err != nil {
			log.Printf("[Register][Err]%s\n", err.Error())
			c.Data["IsLogin"] = false
			c.Data["Error"] = "注册失败：" + err.Error()
			Respjson["IsLogin"] = false
			Respjson["Error"] = "注册失败：" + err.Error()
			if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
				c.Data["IsLogin"] = false
				c.TplName = "bad.html"
			} else {
				c.Data["json"] = Respjson
				c.ServeJSON()
			}
			return
		}

		log.Printf("[Register][Reg Success]username=%s,pwd=%s\n", username, password)
		c.Data["IsLogin"] = true
		c.SetSession("Username", username)
		c.Data["Username"] = c.GetSession("Username")
		Respjson["IsLogin"] = true
		Respjson["Username"] = username
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "index.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	} else {
		c.Data["IsLogin"] = true
		c.Data["Username"] = username
		Respjson["IsLogin"] = true
		Respjson["Username"] = username
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "index.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	}
}

func (c *UserController) PostLogin() {
	username := c.GetSession("Username")
	Respjson := map[string]interface{}{"Server": "AlbumHost Server"}
	if username == nil {
		username := c.GetString("Username") // login.html中传过来的数据
		password := c.GetString("Password")
		log.Printf("[Login][Login Info]username=%s,pwd=%s\n", username, password)

		u := models.User{
			Username: username,
		}
		err := u.ReadByName()
		if err != nil {
			log.Printf("[Login][Err]%s\n", err.Error())
			c.Data["IsLogin"] = false
			c.Data["Error"] = "登录失败：" + err.Error()
			Respjson["IsLogin"] = false
			Respjson["Error"] = "登录失败：" + err.Error()
			if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
				c.TplName = "bad.html"
			} else {
				c.Data["json"] = Respjson
				c.ServeJSON()
			}
			return
		}
		if password != u.Password {
			log.Printf("[Login][Err]%s\n", "Wrong Password")
			c.Data["IsLogin"] = false
			c.Data["Error"] = "登录失败：" + "Wrong Password"
			Respjson["IsLogin"] = false
			Respjson["Error"] = "登录失败：" + "Wrong Password"
			if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
				c.TplName = "bad.html"
			} else {
				c.Data["json"] = Respjson
				c.ServeJSON()
			}
			return
		}
		log.Printf("[Login][Login Success]username=%s,pwd=%s\n", username, password)
		c.SetSession("Username", username)
		c.Data["IsLogin"] = true
		c.Data["Username"] = c.GetSession("Username")
		Respjson["IsLogin"] = true
		Respjson["Username"] = c.GetSession("Username")
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "index.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	} else {
		c.Data["IsLogin"] = true
		c.Data["Username"] = username
		Respjson["IsLogin"] = true
		Respjson["Username"] = username
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "index.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	}
}
