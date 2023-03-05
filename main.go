package main

import (
	"blog/core"
	"blog/global"
	"fmt"
)

func main() {
	//读取配置文件
	core.InitCore()
	fmt.Println(global.Config)
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
