package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{7, 859, 9, 5, 5, 14, 1548, 78, 9}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d\n", nums[i])
	}
}
