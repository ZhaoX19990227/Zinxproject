package znet

import (
	"Zinxproject/zinx/ziface"
	"errors"
	"fmt"
	"net"
)

// IServer接口的实现
type Server struct {
	// 服务器名称
	Name string
	// 服务器ip版本
	IPVersion string
	// 服务器ip
	IP string
	// 服务器端口
	Port int
}

// 定义当前客户端连接的所绑定的hanler api，目前写死 应由客户自定义handler
func CallbackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显功能
	fmt.Println("[Conn Handler] CallBackToClient....")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBack err")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP :%s,Port %d,is starting\n", s.IP, s.Port)
	go func() {
		// 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error", err)
			return
		}
		// 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listener ", s.IPVersion, "error ", err)
			return
		}
		fmt.Println("start zinx server succ", s.Name, "succ listening")

		var cid uint32
		cid = 0

		// 阻塞的等待客户端连接，处理读写
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept conn err", err)
				continue
			}
			//// 已建立连接 处理业务逻辑
			//go func() {
			//	for {
			//		buf := make([]byte, 512)
			//		count, err := conn.Read(buf)
			//		if err != nil {
			//			fmt.Println("receive buf err", err)
			//			continue
			//		}
			//		fmt.Printf("receive client buf %s,count %d\n", buf, count)
			//		// 回显
			//		if _, err := conn.Write(buf[:count]); err != nil {
			//			fmt.Println("write back err", err)
			//			continue
			//		}
			//	}
			//}()

			// 将处理新连接的业务方法和conn进行绑定，得到连接模块
			dealConn := NewConnection(conn, cid, CallbackToClient)
			cid++
			// 启动当前的连接业务处理
			go dealConn.Start()
		}
	}()
}
func (s *Server) Stop() {

}
func (s *Server) Serve() {
	s.Start()

	//TODO:做一些启动服务后的额外业务

	// 阻塞状态，因为Start是异步的
	select {}
}

func NewServer(name string) ziface.IServer {
	server := &Server{
		name,
		"tcp4",
		"0.0.0.0",
		8999,
	}
	return server
}
