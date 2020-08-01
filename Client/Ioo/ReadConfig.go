package Ioo

import (
	"fmt"
	"io"
	"os"
)

func OpFile() string {
	n, _ := os.Getwd()
	file, err := os.Open(n + "\\Config\\Config.txt")
	if err != nil {
		fmt.Println("读取配置文件出现问题!")
	}
	defer file.Close()
	fl := make([]byte, 512)
	var nn int
	for {
		n, err := file.Read(fl)
		if err == io.EOF {
			break
		}
		nn += n
	}
	return string(fl[:nn])
}
