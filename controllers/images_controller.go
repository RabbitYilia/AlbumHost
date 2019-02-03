package controllers

import (
	"AlbumHost/models"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ImageController struct {
	beego.Controller
}

func (c *ImageController) List() {
	username := c.GetSession("Username")
	Respjson := map[string]interface{}{"Server": "AlbumHost Server"}
	if username == nil {
		Respjson["Error"] = "No Login"
		Respjson["IsLogin"] = false
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "login.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	} else {
		u := models.User{
			Username: username.(string),
		}
		err := u.ReadByName()
		if err != nil {
			log.Printf("[File][List][Err]%s\n", err.Error())
			c.Data["IsLogin"] = true
			c.Data["Error"] = err.Error()
			Respjson["IsLogin"] = true
			Respjson["Error"] = err.Error()
			if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
				c.TplName = "bad.html"
			} else {
				c.Data["json"] = Respjson
				c.ServeJSON()
			}
			return
		}

		var Files []*models.File
		o := orm.NewOrm()
		o.QueryTable("file").Filter("Ownuserid", u.Id).All(&Files)

		c.Data["Username"] = username
		c.Data["Files"] = Files
		Respjson["Username"] = username
		Respjson["Files"] = Files
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "list.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
	}
}

func (c *ImageController) Upload() {
	username := c.GetSession("Username")
	Respjson := map[string]interface{}{"Server": "AlbumHost Server"}
	if username == nil {
		Respjson["Error"] = "No Login"
		Respjson["IsLogin"] = false
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.TplName = "login.html"
		} else {
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
		return
	} else {
		u := models.User{
			Username: username.(string),
		}
		err := u.ReadByName()
		if err != nil {
			log.Printf("[Upload][Err]%s\n", err.Error())
			c.Data["Error"] = "上传失败：" + err.Error()
			Respjson["Error"] = "上传失败：" + err.Error()
			if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
				c.Data["IsLogin"] = true
				c.Data["Username"] = username
				c.TplName = "bad.html"
			} else {
				Respjson["IsLogin"] = true
				Respjson["Username"] = username
				c.Data["json"] = Respjson
				c.ServeJSON()
			}
			return
		}

		f, h, _ := c.GetFile("file")
		f.Close()
		c.SaveToFile("file", "./data/"+u.Username+"-"+h.Filename)

		file := models.File{
			Filename:   h.Filename,
			Status:     "Uploaded",
			Path:       "data/" + u.Username + "-" + h.Filename,
			UploadTime: time.Now(),
			Ownuserid:  u.Id,
		}
		err = file.Create()
		if err != nil {
			log.Printf("[Upload][Err]%s\n", err.Error())
			c.Data["Error"] = "上传失败：" + err.Error()
			Respjson["Error"] = "上传失败：" + err.Error()
			if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
				c.Data["IsLogin"] = true
				c.Data["Username"] = username
				c.TplName = "bad.html"
			} else {
				Respjson["IsLogin"] = true
				Respjson["Username"] = username
				c.Data["json"] = Respjson
				c.ServeJSON()
			}
			return
		}
		Respjson["Success"] = "Upload Success"
		if strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") == true {
			c.Data["IsLogin"] = true
			c.Data["Username"] = username
			c.Redirect("/list", 302)
		} else {
			Respjson["IsLogin"] = true
			Respjson["Username"] = username
			c.Data["json"] = Respjson
			c.ServeJSON()
		}
		return
	}
}
