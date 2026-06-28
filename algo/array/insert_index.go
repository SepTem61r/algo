package main

import "fmt"

func insert(nums []int, num int, index int) {
	//遍历数组
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	nums[index] = num
}

func main() {
	nums := []int{1, 2, 3, 4}
	insert(nums, 6, 2)

	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d\n", nums[i])
	}

}
