package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User2 struct {
	Id     int
	Name   string
	Age    int    `gorm:"default:18"`
	Gender string `gorm:"default:男"`
}

func main() {
	db, err := gorm.Open(mysql.Open("root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // 禁用表名复数形式，例如User的表名默认是users,
	}})
	if err != nil {
		panic(err)
	}
	// 定义db.DB(),为了调用Close
	sqldb, _ := db.DB()
	defer sqldb.Close()

	// 自动迁移
	db.AutoMigrate(&User2{})

	// 创建一个实例
	// user := User2{Name: "xiaoming", Age: 1, Gender: "boy"}
	// db.Create(&user)

	// 创建多个实例时，需要[]切片
	// var user1 = []User2{{Name: "xiaoming2"}, {Name: "xiaoming3"}, {Name: "xiaoming4"}}
	// db.Create(&user1)

	// 指定结构体进行创建实例，使用map，key，vault，{}为空，也可以初始化
	// db.Model(&User2{}).Create(map[string]interface{}{
	// 	"Name": "xiaoming5", "Age": 4, "Geder": "girl",
	// })

	// 使用map创建多个实例时，需要map切片
	// db.Model(&User2{}).Create([]map[string]interface{}{
	// 	{"Name": "xiaomign_6", "Age": 6, "Gender": "girl"},
	// 	{"Name": "xiaomign_7", "Age": 7, "Gender": "boy"},
	// })

	// 编写

	// user3 := User2{Name: "xiaowang", Age: 18, Gender: "boy"}
	// db.Create(&user3)
	// fmt.Println(user3)

	// user4 := []User2{
	// 	{Name: "xiaowang1", Age: 18, Gender: "boy"},
	// 	{Name: "xiaowang2", Age: 19, Gender: "boy"}}
	// db.Create(&user4)
	// fmt.Println(user4)

	// 错误使用

	// 您并不总是需要在使用 GORM 时使用 Model() 方法
	// 如果您将结构体值作为参数传递给 CRUD 方法之一
	// GORM 将自动从结构体类型推断出模型
	// user5 := map[string]interface{}{"Name": "laowang3", "Age": 20, "Gender": "boy"}
	// db.Create(&user5)
	// fmt.Println(user5)

	// 使用 Model() 方法指定要操作的模型，并传入一个指向空的 User2 结构体的指针
	// 调用 Create() 方法并传入一个指向包含新记录数据的映射的指针
	// user6 := []map[string]interface{}{
	// 	{"Name": "xiaowang4", "Age": 18, "Gender": "boy"},
	// 	{"Name": "xiaowang5", "Age": 19, "Gender": "boy"},
	// }
	// db.Create(&user6)
	// fmt.Println(user6)

	// 正确使用

	// 用于指定您要操作的模型。它用于为当前操作设置模型，并返回一个新的 *DB 实例，
	// 可以用来对指定的模型执行进一步的操作
	// 传入一个空的 User 结构体指针，以指定我们要操作 User 模型。
	// 然后我们在返回的 *DB 实例上调用 Create() 方法，在数据库中创建一个新记录。
	// db.Model(&User2{}).Create(map[string]interface{}{"Name": "laowang", "Age": 20, "Gender": "girl"})

	// db.Model(&User2{}).Create([]map[string]interface{}{
	// 	{"Name": "xiaowang_6", "Age": 21, "Gender": "boy"},
	// 	{"Name": "xiaowang_7", "Age": 22, "Gender": "boy"},
	// })

}
