package Send

import (
	"fmt"
	"net"
	"Client/Connenct"
)

type Send interface {
	Send()
}

type Sends struct {
	Conn  net.Conn
	State int
}

func (state *Sends) Sendm() {
	defer Connenct.Conn()
	if state.State == 1 {
		var str string
		for {
			fmt.Println("请输入要发送的内容:")
			fmt.Scanln(&str)
			if str == "0.0" {
				return
			}
			_, err := state.Conn.Write([]byte(str))
			if err != nil {
				fmt.Println("发送时出现错误，请不要重试!")
			}
		}
	}
}
