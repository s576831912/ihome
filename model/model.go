package model

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Stu struct {
	gorm.Model
	Name string
	PassWord string
}


var GlobalDB *gorm.DB

func InitModel() {
	//打开数据库  驱动名  连接字符串   返回的是一个数据库连接池
	db,err:=gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/ihome")
	if err != nil{
		fmt.Println("数据库链接失败")
		return
	}
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(30)
	db.DB().SetConnMaxLifetime(1800)
	db.SingularTable(true)
	GlobalDB=db
	//自动迁移 在gorm中建表默认是复数形式
	db.AutoMigrate(new(Stu))
}
