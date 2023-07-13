package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	gorm.Model
	Name         string
	Age          int
	Birthday     *time.Time
	Email        string
	Role         string
	MemberNumber *string
	Num          int
	Adder        string
	IngoreMe     int
}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// 修改将User表名改为 “profiles”
func (Animal) TableName() string {
	return "yong"
}

func (u User) Fils() string {
	if u.Role == "admin" {
		return "admin_users"
	} else {
		return "users"
	}
}

func main() {

	// 禁用默认表名的复数形式，如果为ture，则 ‘User’ 的默认表名是 ‘user’
	db, err := gorm.Open(mysql.Open("root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // 禁用表名复数形式，例如User的表名默认是users,
	}})
	if err != nil {
		panic(err)
	}
	sqldb, _ := db.DB()
	defer sqldb.Close()
	// 函数可以自动迁移您的模式，以保持您的模式最新
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

}
