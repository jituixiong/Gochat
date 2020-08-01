package Ioo

import (
	"fmt"
	"net"
)

func Read(conn net.Conn) {
	defer conn.Close()
	for {
		by := make([]byte, 96)
		_, err := conn.Read(by)
		if err != nil {
			fmt.Println("有客户端退出")
			return
		}
		fmt.Println(string(by))
	}

}
