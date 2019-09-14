package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `orm:"column(id)"`
	Name string `orm:"column(name)"`
}

func main() {
	_ = orm.RegisterDataBase("default", "mysql", "root:wzypwd@tcp(wangzhengyu.cn:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(User))
	dao := orm.NewOrm()
	var result []User
	// select => sql查询
	count, err := dao.Raw("select * from user where 1 = ?", 1).QueryRows(&result)
	if err == nil {
		fmt.Printf("查询到%d条数据:\n", count)
		fmt.Println(result)
	}
	result = []User{}
	// select => api方式
	count, err = dao.QueryTable("user").Exclude("id", -2).Limit(10, 0).All(&result)
	if err == nil {
		fmt.Printf("查询到%d条数据:\n", count)
		fmt.Println(result)
	}

	// update => sql方式
	sqlResult, err := dao.Raw("update user set name = ? where id = ?", "new_name", 1).Exec()
	if err == nil {
		affected, _ := sqlResult.RowsAffected()
		fmt.Printf("修改了%d行\n", affected)
	}

	// delete => sql方式
	sqlResult, err = dao.Raw("delete from user where name = ?", "new_name").Exec()
	if err == nil {
		affected, _ := sqlResult.RowsAffected()
		fmt.Printf("删除了%d行\n", affected)
	}

	// insert => sql方式
	var users []User = []User{
		{1, "newName"},
	}
	// bulk表示并行数量
	count, err = dao.InsertMulti(1, &users)
	if err == nil {
		fmt.Printf("插入了%d行\n", count)
	} else {
		fmt.Println(err)
	}
}
