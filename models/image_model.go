package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

type Image struct {
	Id         int `orm:"column(id);auto;pk" description:"id"` // 主键
	Filename   string
	Sha256hash string
	Status     string
	Path       string
	Ownuserid  int
	UploadTime time.Time
}

func (f *Image) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(f)
	if err == nil {
		log.Printf("Create Image:%s success!\n", f.Filename)
	} else {
		log.Printf("Create Image:%s failed!,%s\n", f.Filename, err.Error())
	}
	return err
}

func (f *Image) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(f)
	if err == nil {
		log.Printf("Update Image:%s success!\n", f.Filename)
	} else {
		log.Printf("Update Image:%s failed!,%s\n", f.Filename, err.Error())
	}
	return err
}
func init() {
	orm.RegisterModel(new(Image))
}
