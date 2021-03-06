package models

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接初始化
func Init() {
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		appConf.String("database::db_user"),
		appConf.String("database::db_passwd"),
		appConf.String("database::db_host"),
		appConf.String("database::db_port"),
		appConf.String("database::db_name"),
		appConf.String("database::db_charset"))

	// 设置数据库基本信息，注册default 别名alias
	orm.RegisterDataBase("default", "mysql", conn)

	//    自动建表
	// 别名
	name := "default"
	// 强制更新
	force := false
	// 是否可见创建过程
	verbose := true
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	orm.RunCommand()
}

//返回带前缀的表名
func TableName(str string) string {
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	return appConf.String("database::db_prefix") + str
}

// BaseQueryParam 用于查询的类
type BaseQueryParam struct {
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
}