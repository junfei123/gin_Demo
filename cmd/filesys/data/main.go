package main

import (
	"ginDemo/dao"
	"ginDemo/router"

	"github.com/glebarez/sqlite" //不依赖cgo
	//"gorm.io/driver/sqlite" // 如果需要使用cgo，可以改用
	"gorm.io/gorm"
)

func main() {
	//读取db文件。
	db, err := gorm.Open(sqlite.Open("../../filesys.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dao.SetDefault(db) // 让 gen 生成的代码使用这个 db
	//路由
	r := router.InitRouter()

	r.Run(":8080")

}
