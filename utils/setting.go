package utils

import (
	"fmt"

	"gopkg.in/ini.v1"  // goini
)

// go语言中，变量名以及函数名首字母大写，则外部可以访问
var (  // 声明全局变量名称以及变量类型，使用goini读取配置参数，之后直接调用即可
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPost     string
	DbUser     string
	DbPassWord string
	DbName     string
)

// 包初始化时会运行init函数
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误", err)
	}
	LoadServer(file)  // 读取server分区
	LoadDb(file)  // 读取db分区
}

// 加载server分区参数
func LoadServer(file *ini.File) {
	// 从分区对应关键字取值，以及默认值
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpProt").MustString(":8080")
	JwtKey = file.Section("server").Key("JwtKey").MustString("adajdafaf")
}

// 加载database分区参数
func LoadDb(file *ini.File) {
	// 从分区对应关键字取值，以及默认值
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPost = file.Section("database").Key("DbPost").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("admin")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("autotrans")
}
