package main

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 模型是一种特殊的结构体，用于定义数据库表中的数据结构
type User struct {
	ID int
	// 模型字段类型为指针意味着该字段可以为 NULL // 字段没写就是没传该字段（没写入库），使用了指针，可以传入字段内容为空，写入数据库。
	// Name *string `gorm:"default:'小笨蛋'"`
	Name sql.NullString `gorm:"default:'小笨蛋'"` // sql. string有值就用，没值则使用默认
	Age  *int
}

func main() {
	// engin := gin.Default()
	// 禁用默认表名的复数形式，如果为ture，则 ‘User’ 的默认表名是 ‘user’
	db, err := gorm.Open(mysql.Open("root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // 禁用表名复数形式，例如User的表名默认是users,
	}})
	if err != nil {
		panic(err)
	}
	sqldb, _ := db.DB()
	defer sqldb.Close()
	// g1 := engin.Group("/api")
	// if g1 != nil {

	// }

	// 2, 自动迁移数据
	db.AutoMigrate(&User{})

	// 3,创建一条数据
	// 1结构体类型为指针,想传入空，又不使用默认值
	// u := &User{Name: new(string), Age: 48} // new(string)为string指针
	// 2,sql中的东西
	age := 19
	u := &User{Name: sql.NullString{String: "", Valid: true}, Age: &age}
	db.Debug().Create(&u)
}
