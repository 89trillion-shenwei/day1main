package util

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

//解析出端口号
func AnalyIni() (s string) {
	c, err1 := goconfig.LoadConfigFile("../conf/app.ini") //读取配置文件
	if err1 != nil {
		fmt.Println("读取app.ini时出现错误")
		fmt.Println(err1)
		return
	}
	port, err2 := c.GetValue("server", "HttpPort")
	if err2 != nil {
		fmt.Println("获取端口号时出现错误")
		fmt.Println(err2)
		return
	}
	return port

}
