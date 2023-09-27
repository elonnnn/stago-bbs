package main

import (
	"fmt"

	"github.com/elonnnn/stago-bbs-service/core"
	"github.com/elonnnn/stago-bbs-service/global"
)

func main() {
	fmt.Println("-------------------Welcome to stago-admin service-----------------")
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	// fmt.Println(global.GVA_VP.Get("mysql.port"))

}
