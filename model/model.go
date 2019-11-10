package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Stu struct {
	Id int
	Name string
	PassWord string
}


var GlobalDB *gorm.DB

func InitModel() {
	//打开数据库    参数1：驱动名  参数2：链接字符串
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ihome")
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}

	//连接池设置
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(30)
	db.DB().SetConnMaxLifetime(60 * 30)
//设置表明为单数形式
db.SingularTable(true)
	GlobalDB = db
	db.AutoMigrate(new(Stu))
}
