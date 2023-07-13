package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	Id     int
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open("root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	//创建数据表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	// u1 := &UserInfo{Id: 1, Name: "yong", Gender: "nan", Hobby: "lanqiu"}
	// db.Create(&u1)

	// 查询第一条数据
	var u UserInfo
	db.First(&u) // db.first()就是查询第一条，想要更改数据必需用指针，将查询到的数据保存到 u 中！
	fmt.Printf("u:%#v\n", &u)

	// 更新数据
	db.Model(&u).Update("hobby", "双色球")

	// 删除
	db.Delete(&u)
}
