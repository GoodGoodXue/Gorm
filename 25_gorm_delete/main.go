package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User2 struct {
	Id     int
	Name   string
	Age    int
	Gender string
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

	db.AutoMigrate(&User2{})

	// 定义User2 类型结构体
	// var user User2

	// 结构体创建，可以批量插入
	// var user = User2{Name: "two", Age: 19, Gender: "boy"}
	// db.Create(&user)

	// db.Debug().Create(&User2{Name: "three", Age: 23, Gender: "girl"})

	// 使用map 批量插入
	// db.Model(&user).Create([]map[string]interface{}{
	// 	{"Name": "wang", "Age": 18, "Gender": "girl"},
	// 	{"Name": "sam", "Age": 10, "Gender": "boy"}})

	// 需要指定主键删除
	// user.Id = 2
	// db.Delete(&user)

	// 条件删除
	// db.Debug().Where("name =?", "xiaowang2").Delete(&user)

	// 根据主键删除
	// db.Debug().Delete(&user, 4)
	// db.Debug().Delete(&user, []int{5, 6})

	// 批量删除 匹配删除
	// db.Where("Name LIKE ?", "%two%").Delete(&user)
	// db.Delete(&user, "name LIKE ?", "%three%")

	// 编写

	// 须指定主键删除，否则触发全局删除
	var user User2
	user.Id = 1
	db.Delete(&user)

	// 根据主键删除
	db.Delete(&user, 3)
	db.Delete(&user, []int{1, 2})

	// 批量 匹配 条件删除
	db.Where("name = ?", "wang").Delete(&user)

	db.Where("name LIKE ?", "%san%").Delete(&user)

}
