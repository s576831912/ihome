package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Stu struct {
	gorm.Model
	Name     string
	PassWord string
}

var GlobalDB *gorm.DB

func InitModel() {
	//打开数据库    参数1：驱动名  参数2：链接字符串
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ihome?parseTime=true")
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

//插入数据函数
func InsertData() {
	//有一个赋完值结构体对象
	var stu Stu
	stu.Name = "bj5q"
	stu.PassWord = "123456"
	err := GlobalDB.Create(&stu).Error
	if err != nil {
		fmt.Println("创建数据失败")
		return
	}
	fmt.Println(stu)
}

//查找数据
func SearchData() {
	var stu Stu
	stu.ID = 1
	if err := GlobalDB.Where("name = ?", "itcast").Where("pass_word = ?", "00000").First(&stu).Error; err != nil {
		fmt.Println("查询错误，err:", err)
		return
	}
	fmt.Println(stu)
}

//更新数据
func UpdateData() {
	var stu Stu
	//stu.ID=1
	stu.Name = "itcast"
	stu.PassWord = "111222"
	//正常更新
	/*if err:=GlobalDB.Save(&stu).Error;err!=nil{
		fmt.Println("更新数据失败,err:",err)
		return
	}*/
	//按照条件更新
	err := GlobalDB.Model(&stu).Where("name = ?", "itcast").Update("pass_word", "00000").Error
	if err != nil {
		fmt.Println("更新密码失败,err:", err)
		return
	}
	fmt.Println(stu)
}

//删除数据
func DeleteData() {
	var stu Stu

	if err:=GlobalDB.Where("id=1").Delete(&stu).Error;err!=nil{
		fmt.Println("删除数据错误,err:",err)
		return
	}
	fmt.Println(stu)
}
