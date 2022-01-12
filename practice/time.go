package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type TestDatetime struct {
	Id int
	DatetimeAt time.Time
}

func main() {
	time.ParseInLocation()
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dest := make([]TestDatetime, 0)
	res := db.Find(&dest)
	//db.Save(&TestDatetime{
	//	Id:         0,
	//	DatetimeAt: time.Now(),
	//})
	fmt.Println(res, err)
	fmt.Println(dest)
	//db, err :=
	//if err != nil {
	//	panic(err)
	//}
	//// See "Important settings" section.
	//db.SetConnMaxLifetime(time.Minute * 3)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
}
