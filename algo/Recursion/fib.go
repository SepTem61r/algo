package main

import "fmt"

/*
我们在函数内递归调用了两个函数
这意味着从一个调用产生了两个调用分支
这样不断递归调用下去，最终将产生一棵层数为n的递归树
*/
func fib(n int) int {

	if n == 1 || n == 2 {
		return n - 1
	}
	
	res := fib(n-1) + fib(n-2)
	return res
}

func main() {
	fmt.Printf("%d", fib(3))
}
