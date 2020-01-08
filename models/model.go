package models

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

type SalEmp struct {
	Name string
	PayByQuarter []int

}

func init() {
	db, _ := gorm.Open("postgres", "host=localhost user=postgres dbname=BBS sslmode=disable password=loveys1314")
	fmt.Println(db,"--Acc----")
	db.SingularTable(true)
	//db.AutoMigrate(&Diary{})

	//var   diary Diary
	//
	//db.First(&diary)
	//fmt.Println(diary.Photos,"-----attay------")

	var   diary SalEmp

	db.First(&diary)
	fmt.Println(diary.PayByQuarter,"-----attay------")

	defer db.Close()
}


var DB *gorm.DB
var Client *redis.Client

// 初始化数据库连接等
func Initialized() {
	// 申明变量
	var (
		err error
	)
	// 连接数据库
	DB, err = gorm.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatalf("database connection error: %v", err)
		return
	}

	Client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	DB.SingularTable(true)
	//DB.AutoMigrate(&Classify{}, &Diary{})


	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(20000)
	// debug 模式开启sql日志
	DB.LogMode(true)
}

// 关闭数据库连接
func CloseDb() {
	DB.Close()
}
