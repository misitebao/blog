package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	DB *gorm.DB
)
func init()  {
	// 连接数据库
	DB, err := gorm.Open("mysql", "root:123456@/mtblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	log.Print("连接数据库成功")

	defer DB.Close()
}
