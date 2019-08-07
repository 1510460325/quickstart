package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

var (
	dao orm.Ormer
)

func init() {
	// set default database
	_ = orm.RegisterDataBase("default", "mysql", "root:wzypwd@tcp(wangzhengyu.cn:3306)/test?charset=utf8", 30)
	// register model
	orm.RegisterModel(new(User))
	// create table
	_ = orm.RunSyncdb("default", false, true)
	dao = orm.NewOrm()
	fmt.Println("db init completed.")
}

func Insert(user User) int64 {
	id, err := dao.Insert(&user)
	if err == nil {
		return id
	} else {
		return -1
	}
}

func Update(user User) int64 {
	num, err := dao.Update(&user)
	if err == nil {
		return num
	} else {
		return -1
	}
}

func Read(id int) User {
	query := User{Id: id}
	err := dao.Read(&query)
	if err == nil {
		return query
	} else {
		fmt.Printf("query is wrong, error is %v \n", err)
		return User{Id: -1}
	}
}

func Delete(id int) int64 {
	query := User{Id: id}
	num, err := dao.Delete(&query)
	if err == nil {
		return num
	} else {
		return -1
	}
}
