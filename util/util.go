package util

import (
	"fmt"
)

// 操作提示
func OpTips(tips []string) {
	for k, v := range tips {
		fmt.Printf("(%d) : %s\n", k, v)
	}
	println("退出请输 x")
}

func TipStart()  {
	fmt.Println("=============== start =============== >>>>>>>>>>")
}
func TipEnd()  {
	fmt.Println("=============== end   =============== >>>>>>>>>>")
}
