package main

import (
	"Zinxproject/zinx/ziface"
	"Zinxproject/zinx/znet"
	"fmt"
)

type PingRouter struct {
	znet.BaseRouter
}

// Test PreHandler
func (this *PingRouter) PreHandler(request ziface.IRequest) {
	fmt.Println("PreHandler")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("before ping..."))
	if err != nil {
		fmt.Println("before ping error", err)
		return
	}

}

// Test Handler
func (this *PingRouter) Handler(request ziface.IRequest) {
	fmt.Println("Handler")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("ping..."))
	if err != nil {
		fmt.Println(" ping error", err)
		return
	}
}

// Test PostHandler
func (this *PingRouter) PostHandler(request ziface.IRequest) {
	fmt.Println("PostHandler")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("post ping..."))
	if err != nil {
		fmt.Println("post ping error", err)
		return
	}
}

func main() {
	server := znet.NewServer("[zinx V0.2]")
	// 给当前zinx框架添加一个自定义的router
	server.AddRouter(&PingRouter{})
	server.Serve()
}
