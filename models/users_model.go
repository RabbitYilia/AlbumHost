package models

import (
	"log"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int `orm:"column(id);auto" description:"id"` // 主键
	Username string
	Password string
	Email    string
}

func (u *User) ReadByName() (err error) {
	o := orm.NewOrm()
	err = o.Read(u, "Username")
	if err == nil {
		log.Printf("Query User:%s success!\n", u.Username)
	} else {
		log.Printf("Query User:%s failed!,%s\n", u.Username, err.Error())
	}
	return err
}

func (u *User) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(u)
	if err == nil {
		log.Printf("Create User:%s success!\n", u.Username)
	} else {
		log.Printf("Create User:%s failed!,%s\n", u.Username, err.Error())
	}
	return err
}

func (u *User) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(u)
	if err == nil {
		log.Printf("Update User:%s success!\n", u.Username)
	} else {
		log.Printf("Update User:%s failed!,%s\n", u.Username, err.Error())
	}
	return err
}
func init() {
	orm.RegisterModel(new(User))
}
