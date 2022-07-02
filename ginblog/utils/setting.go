package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}
	LoadServer(file)
	LoadData(file)

}
func LoadServer(file *ini.File) {
	// 在server这个section下面的AppMode取不到值时，取默认值debug
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("postgres")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("5432")
	DbUser = file.Section("database").Key("DbUser").MustString("postgres")
	DbPassword = file.Section("database").Key("DbPassword").MustString("u5398681234QWer")
	DbName = file.Section("database").Key("DbName").MustString("gin_blog")
}
