package main

import "fmt"

func delete(nums []int, index int) {
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	delete(nums, 1)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d\n", nums[i])
	}
}
