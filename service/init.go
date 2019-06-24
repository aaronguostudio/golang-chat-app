package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// init 函数在项目初始化的时候会自动运行

var DbEngin *xorm.Engine

func init () {
	drivername := "mysql"
	datasourcename := "root:root@(127.0.0.1:3306)/chat?chartset=utf8"
	DbEngin, err := xorm.NewEngine(drivername, datasourcename)
	if err != nil {
		log.Fatal(err.Error())
	}

	DbEngin.ShowSQL(true)
	DbEngin.SetMaxOpenConns(2)

	// 自动设置 User 表
	//DbEngin.Sync2(new (User))
	fmt.Println("DB success")
}