package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var i input

func init() {
	i = input{
		reader: bufio.NewReader(os.Stdin),
	}
}

type input struct {
	reader   *bufio.Reader
	text     string
	autoQuit bool
}

func AutoQuit() {
	i.autoQuit = true
}

func GetInput(tips ...string) string {
	for _, tip := range tips {
		print(tip, " : ")
	}
	var err error
	i.text, err = i.reader.ReadString('\n')
	if err != nil {
		log.Fatalln("从标准输入读取出错", err)
	}
	i.text = strings.TrimRight(i.text, string('\n'))
	if i.autoQuit && ShouldQuit() {
		os.Exit(0)
	}
	return i.text
}

// 退出
func ShouldQuit() bool {
	if i.text == "x" {
		return true
	}

	return false
}
