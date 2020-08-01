package RunServer

import (
	"Gochat/Server/Ioo"
	"fmt"
	"net"
)

func Run() {
	fmt.Println("服务端开始启动!")
	list, err := net.Listen("tcp", "0.0.0.0:12388")
	defer list.Close()
	if err != nil {
		fmt.Println("服务端启动失败！")
	}
	for {
		conn, err := list.Accept()
		if err != nil {
			fmt.Println("服务端启动失败")
			return
		}
		go Ioo.Read(conn)
	}
}
