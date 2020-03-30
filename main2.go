package main

import "fmt"

func main() {
	sum := 0
	for {
		sum++ // 无限循环下去
		fmt.Println(sum)
		if sum == 10000000 {
			break
		}
	}
	fmt.Println(sum) // 无法输出
}
