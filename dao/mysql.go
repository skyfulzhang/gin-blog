package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

// InitDb 数据库初始函数
func InitMysql() {
	sqlStr := "@tcp(127.0.0.1:3306)/study?charset=utf8&parseTime=true&loc=Local"
	db, err = gorm.Open("mysql", sqlStr)
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数！", err)
		panic("failed to connect database")
	}
	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)

	// 自动迁移
	//db.AutoMigrate(&User{}, &Category{}, &Article{})

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)
}

func GetDB() *gorm.DB {
	return db
}
