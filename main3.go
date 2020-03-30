package main

import "fmt"

var by = []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
var yh = []int{7, 9, 3, 4}
var n = []int{4}

func main() {

	fmt.Println(by)

	fmt.Println(yh)

	isyh(n)

}

func isyh(n []int) {

	for i := 0; i < len(yh); i++ {
		if n[0] == yh[i] {
			fmt.Println("是要害")
			break
		}
		fmt.Println("不是")
	}
}
