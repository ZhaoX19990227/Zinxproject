package znet

import (
	"Zinxproject/zinx/ziface"
	"fmt"
	"net"
)

type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn

	// 连接的Id
	ConnID uint32

	// 当前连接状态
	isClosed bool

	// 当前连接所绑定的处理业务方法API
	handleAPI ziface.HandleFunc

	// 告知当前连接停止 channel
	ExitChan chan bool
}

// 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is Running...")
	defer fmt.Println("connID=", c.ConnID, "Reader is exit,remote addr is=", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中
		buf := make([]byte, 512)
		// 读到buf中
		count, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Conn read err", err)
			continue
		}
		// 调用当前连接绑定的HandlerAPi
		if err := c.handleAPI(c.Conn, buf, count); err != nil {
			fmt.Println("ConnID=", c.ConnID, "handler error ", err)
			break
		}
	}
}

// 启动连接
func (c *Connection) Start() {
	fmt.Println("conn start(),connID = ", c.ConnID)
	//TODO 启动从当前连接的读数据的业务
	go c.StartReader()

}

// 停止连接
func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID = ", c.ConnID)
	if c.isClosed {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

// 获取当前连接的socket conn
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前连接的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的TCP状态和Ip端口
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send(data []byte) error {
	return nil
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	connection := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callback_api,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return connection
}
