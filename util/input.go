package util

import (
	"bufio"
	"errors"
	"github.com/astaxie/beego/logs"
	"log"
	"os"
	"strconv"
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
	bufText  []string
	autoQuit bool
}

// 可单元测试用
func SetInput(lines ...string) {
	for _, line := range lines {
		i.bufText = append(i.bufText, line)
	}
}

func AutoQuit() {
	i.autoQuit = true
}

func GetInput(tips ...string) string {
	for _, tip := range tips {
		print(tip, " : ")
	}
	var err error
	err = readIntoText()
	if err != nil {
		log.Fatalln("从标准输入读取出错", err)
	}
	i.text = strings.TrimRight(i.text, string('\n'))
	if i.autoQuit && ShouldQuit() {
		os.Exit(0)
	}
	return i.text
}

func GetChoice(tips ...string) (choice int, err error) {
	choiceStr := GetInput(tips...)
	choice, err = strconv.Atoi(choiceStr)
	if err != nil {
		logs.Error("输入选项获取失败", err)
		err = errors.New("输入选项获取失败")
		return
	}
	return
}

func readIntoText() (err error) {
	if len(i.bufText) > 0 {
		i.text = i.bufText[0]
		i.bufText = i.bufText[1:]
		println(i.text)
	} else {
		i.text, err = i.reader.ReadString('\n')
	}
	return
}

// 退出
func ShouldQuit() bool {
	if i.text == "x" {
		return true
	}

	return false
}
