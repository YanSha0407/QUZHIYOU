package models


import (
	"fmt"
	"github.com/jinzhu/gorm"
     "log"
)

var DB *gorm.DB



// 初始化数据库连接等
func Initialized() {
	// 申明变量
	var (
		err error
		//connectSql = "root:loveys1314@tcp(127.0.0.1:3306)/qzy_official_service?charset=utf8&parseTime=True&loc=Local"
		connectSql = "root:123qaz!@#@tcp(39.97.230.148:3306)/qzy_official_service?charset=utf8&parseTime=True&loc=Local"
		databaseCfg = "mysql"
	)
	// 连接数据库
	DB, err = gorm.Open(databaseCfg, connectSql)
	if err != nil {
		log.Fatalf("database connection error: %v", err);
		return
	}




	DB.SingularTable(true)
	fmt.Println("数据库连接成功")
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	// debug 模式开启sql日志
	DB.LogMode(true)
}

// 关闭数据库连接
func CloseDb() {
	DB.Close()
}
