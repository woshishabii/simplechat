package main

import (
	"github.com/CAP/simplechat/conf"
	"github.com/CAP/simplechat/server"
)

func main() {
	// 初始化配置文件
	conf.Init()
	// 获取Router
	r := server.NewRouter()
	// 在指定的端口上运行
	r.Run(":8192")
}
