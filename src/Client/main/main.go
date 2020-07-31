package main

import (
	"Client/Connenct"
	"Client/Send"
)

func main() {
	co, err := Connenct.Conn()
	c := Send.Sends{co, err}
	c.Sendm()
}