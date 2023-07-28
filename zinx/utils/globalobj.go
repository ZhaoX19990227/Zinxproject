package utils

import (
	"Zinxproject/zinx/ziface"
	"encoding/json"
	"os"
)

// 全局参数
type GlobalObj struct {
	// server
	TcpServer ziface.IServer //当前zinx全局的server对象
	Host      string
	TcpPort   int
	Name      string

	// Zinx
	Version        string
	MaxConn        int    // 最大连接数
	MaxPackageSize uint32 // 当前Zinx框架数据包的最大值
}

func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	// 将data映射到GlobalObject中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

/*
*
定义一个全局的对外GlobalObj对象
*/
var GlobalObject *GlobalObj

// 用来初始化GlobalObject
func init() {
	// 默认值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "V0.3",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	// 尝试从conf/zinx.json文件中加载用户自定义的参数
	GlobalObject.Reload()
}
