package main

import (
	"Gochat/Client/Connenct"
	"Gochat/Client/Send"
)

func main() {
	co, err := Connenct.Conn()
	c := Send.Sends{co, err}
	c.Sendm()
}
