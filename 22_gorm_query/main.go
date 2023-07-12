package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User1 struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	db, err := gorm.Open(mysql.Open("root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // 禁用表名复数形式，例如User的表名默认是users,
	}})
	if err != nil {
		panic(err)
	}

	sqldb, _ := db.DB()
	defer sqldb.Close()

	db.AutoMigrate(&User1{})

	// db.Create(&User1{Name: "xiaowang", Age: 18})
	// db.Create(&User1{Name: "laowang", Age: 28})
	user := []*User1{&User1{Name: "小1", Age: 1}, &User1{Name: "小2", Age: 2}}
	db.Create(&user)
	// 查询
	// var user User1 // 声明模型结构体类型变量user
	// user := new(User1) // new返回指针类型，所以user为指针

	// db.First(user)
	// fmt.Printf("user:%#v\n", user)

	// var users []User1
	// db.Debug().Find(&users)
	// fmt.Printf("user:%#v\n", users)

}
