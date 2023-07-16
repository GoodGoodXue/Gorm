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

	// 定义user为User2的结构体类型
	// var user User2

	// db.First(&user)
	// user.Name = "小明"
	// user.Age = 28

	// Save保存所有字段
	// db.Debug().Save(&user)

	// 使用where需要使用model指定结构体，Where（条件内容）。Update为更新单个字段
	// db.Debug().Model(&user).Where("name = ?", "小明").Update("name", "wanger")

	// db.Debug().Model(&user).Update("Name", "小王")

	// 使用map更新多字段，要是用Updates更新多字段
	// db.Debug().Model(&user).Updates(map[string]interface{}{"name": "laowang", "age": 19, "gender": "girl"})

	// updates更新多字段使用
	// db.Debug().Model(&user).Updates(User2{Name: "wankang", Age: 8})
	// db.Debug().Model(&user).Updates(User2{Name: "", Age: 0, Gender: "false"})

	// 初始化map
	// m1 := map[string]interface{}{
	// 	"name":   "quan",
	// 	"age":    18,
	// 	"gender": "boy",
	// }

	// 更新m1所有接受的字段
	// db.Model(&user).Updates(m1)

	// Select 只更新指定字段
	// db.Debug().Model(&user).Select("name").Updates(m1)

	// Omit 不更新指定字段
	// db.Debug().Model(&user).Omit("age").Updates(m1)

	// gorm.Expr()使用SQL语句表达式
	// db.Model(&User2{}).Where("age").Update("age", gorm.Expr("age +?", 2))

	// 编写
	var user User2
	// 任何其他操作都在查找之后，只有查询到该数据才可进行操作
	db.First(&user)

	// user.Name = "wanger"
	// user.Age = 18
	// user.Gender = "girl"
	// db.Save(&user)

	// 使用更新需要Model，更新单个字段内容（单列）
	db.Model(&user).Update("name", "fanfan")

	// Where添加条件，更新单列
	// db.Debug().Model(&user).Where("name = ?", "wanger").Update("name", "baga")

	// 根据struct，更新多列，只会更新非零值的字段
	db.Model(&user).Updates(User2{Name: "wan", Age: 0, Gender: "girl"})

	// updates，使用map进行批量更新
	db.Debug().Model(&user).Where("name = ?", "wanger").Updates(map[string]interface{}{"name": "wang", "age": 11})

	// select 更新选定字段
	// db.Debug().Model(&user).Select("name").Updates(map[string]interface{}{"name": "wang_1", "age": 12, "gender": "girl"})

	// Omit 更新非选定字段
	// db.Debug().Model(&user).Omit("name").Updates(map[string]interface{}{"name": "wang_1", "age": 18, "gender": "boy"})

	// gorm.Expr 使用SQL表达式更新
	// db.Debug().Model(&user).Where("age").Update("age", gorm.Expr("age + ?", 2))

}
