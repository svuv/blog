package main

import (
	"blog/core"
	"blog/global"
	"fmt"
)

func main() {
	core.InitCore()
	fmt.Println(global.Config)
}
