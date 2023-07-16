package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User1 struct {
	Id   int
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

	// db.Create(&User1{Name: "xiaowang", Age: 38})
	// db.Create(&User1{Name: "laowang", Age: 48})

	// 批量插入
	// user := []*User1{&User1{Name: "小1", Age: 1}, &User1{Name: "小2", Age: 2}}
	// db.Create(&user)

	// 查询 Usre1为类型，user只能存一个支，而切片[]User1为切片类型，可以多个值
	var user []User1 // 声明模型结构体类型变量user

	// user := new(User1) // new返回指针类型，所以user为指针

	// 根据主键查询第一个值
	// db.First(&user)
	// fmt.Printf("user:%#v\n", user)

	// 非主键查询第一个值
	// db.Take(&user)
	// fmt.Printf("user:%#v\n", user)

	// 查询最后一个值
	// db.Last(&user)
	// fmt.Printf("user:%#v\n", user)

	// 定义为map【键类型】外面为值类型，最后的{} 表示创建一个空映射，不包含任何键值对，
	// 创建映射时初始化一些键值对
	// result := map[string]interface{}{
	// 	"key1": "value1",
	// 	"key2": 123,
	// 	"key3": true,
	// }

	// result := map[string]interface{}{}
	// Model为选择结构体进行后续操作
	// db.Model(&User1{}).First(&result)
	// 初始化输出%v 默认输出
	// fmt.Printf("result:%v\n", result)

	// 选择模型查找
	// db.Model(&User1{}).Take(&result)
	// fmt.Printf("result:%v\n", result)

	// 主键查询第一个
	// db.First(&user, 1)

	// 查询全部
	// db.Find(&user)

	// 查询第一二值
	// db.Debug().Find(&user, []int{1, 2})
	// fmt.Println(user)

	// var users []User1
	// db.Debug().Find(&users)
	// fmt.Printf("user:%#v\n", users)

	// 条件查询
	// db.Where("name=?", "xiaowang").First(&user)
	// fmt.Println(user)

	// 查询非xiaowang 所有
	// db.Where("name <> ?", "xiaowang").Find(&user)
	// fmt.Println(user)

	// 查询多个匹配值
	// db.Where("name IN ?", []string{"xiaowang", "laowang"}).Find(&user)
	// fmt.Println(user)

	// 模糊查询
	// db.Where("name LIKE ?", "%xiao%").Find(&user)
	// fmt.Println(user)

	// 匹配查询
	db.Where(&User1{Name: "xiaowang", Age: 18}).First(&user)
	fmt.Println(user)

	// 匹配查询
	// db.Where(map[string]interface{}{"name": "laowang", "age": 28}).Find(&user)
	// db.Where([]int64{1, 3, 5}).Find(&user)
	// fmt.Println(user)

	// 条件选择查询内容
	// db.Select("name", "age").Find(&user)
	db.Select([]string{"name", "age"}).Find(&user)
	fmt.Println(user)

	// var user []User1
	db.First(&user, 1)
	db.Find(&user, []int{1, 2, 3})
	db.Take(&user)
	db.Where("name = ?", "xiaowang").First(&user)
	db.Where("name <> ?", "xiaowang").Find(&user)
	db.Where("name IN ?", []string{"xiaowang", "laowang"}).Find(&user)
	db.Where("name Like ?", "%xiao%").Find(&user)

	db.Where(&User1{Name: "xiaowang", Age: 18}).First(&user)
	db.Where(map[string]interface{}{"name": "xiaowang", "age": 18}).Find(&user)
}
