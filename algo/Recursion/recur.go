package main

import "fmt"

// 正常递归
func recur(n int) int {
	if n == 1 {
		return 1
	}

	res := recur(n - 1)

	return res + n
}

func main() {
	fmt.Printf("%d", recur(10))
}
