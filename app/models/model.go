package models

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB
var Client *redis.Client


// 初始化数据库连接等
func Initialized() {
	// 申明变量
	var (
		err error
		connectSql = "root:loveys1314@tcp(127.0.0.1:3306)/qzy_official_service?charset=utf8&parseTime=True&loc=Local"
		//connectSql = "root:123qaz!@#@tcp(39.97.230.148:3306)/qzy_official_service?charset=utf8&parseTime=True&loc=Local"
		databaseCfg = "mysql"
	)
	// 连接数据库
	DB, err = gorm.Open(databaseCfg, connectSql)
	if err != nil {
		log.Fatalf("database connection error: %v", err);
		return
	}

	Client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(Client,"------")


	DB.SingularTable(true)

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(5000)
	// debug 模式开启sql日志
	DB.LogMode(false)
}

// 关闭数据库连接
func CloseDb() {
	DB.Close()
}
