package Connenct

import (
	"Client/Ioo"
	"fmt"
	"net"
)

func Conn() (net.Conn, int) {
	var state int //连接状态码,0表示失败,1表示成功
	ip := Ioo.OpFile()
	fmt.Println("服务器ip为:" + ip)
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println("连接到服务器时出现错误!")
		state = 0 //连接出现错误状态码赋值为0
		conn = nil
	} else {
		fmt.Println("连接到服务器成功!")
		state = 1 //连接成功状态码赋值为1
	}
	return conn, state
}
