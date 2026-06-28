package main

import "fmt"

/*
尾递归
递归调用是函数返回前的最后一个操作，
这意味着函数返回到上一层级后，无须继续执行其他操作，
因此系统无须保存上一层函数的上下文。
*/
func tailRecur(n, res int) int {
	if n == 0 {
		return res
	}

	return tailRecur(n-1, res+n)
}

func main() {
	fmt.Printf("%d", tailRecur(10, 0))
}
