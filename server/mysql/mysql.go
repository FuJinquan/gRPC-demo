//包名与文件夹名保持一致！
package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Note struct {
	gorm.Model
	Title string
	Text  string
}

//连接mysql数据库
func SetUp() {
	var err error
	db, err = gorm.Open("mysql", "root:950120fjq@/note?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.SingularTable(true)
	//defer db.Close()
	if err := db.AutoMigrate(&Note{}).Error; err != nil {
		panic("创建数据库失败！")
	}
}

//返回db变量
func GetDB() *gorm.DB {
	return db
}
