package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start")
	time.Sleep(1 * time.Second)
	conn, _ := net.Dial("tcp", "127.0.0.1:8999")
	for {
		conn.Write([]byte("hello Zinx V0.3"))
		buf := make([]byte, 512)
		n, _ := conn.Read(buf)
		fmt.Printf("server callback:%s,count = %d\n", buf, n)
		//cpu 阻塞
		time.Sleep(1 * time.Second)
	}
}
