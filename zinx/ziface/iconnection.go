package ziface

import "net"

// 定义连接模块的抽象层

type IConnection interface {
	// 启动连接
	Start()

	// 停止连接
	Stop()

	// 获取当前连接的socket conn
	GetTcpConnection() *net.TCPConn

	// 获取当前连接的连接ID
	GetConnID() uint32

	//获取远程客户端的TCP状态和Ip端口
	RemoteAddr() net.Addr

	// 发送数据
	Send(data []byte) error
}

// 定义处理连接业务的方法  函数类型 三个形参 返回值error
type HandleFunc func(*net.TCPConn, []byte, int) error
